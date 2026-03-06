package processing

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"payment-simulator/internal/cache"
	"payment-simulator/internal/db"
	"payment-simulator/internal/mapping"
	"payment-simulator/internal/models"
	"payment-simulator/internal/routing/outgoing"
	"strings"
	"time"

	"github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.Background()

func init() {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			if oldSnapshotKeys, err := hasOldSnapshots(); err != nil {
				log.Println("There was an error in checking for Old PACS002 Snapshots:", err)
			} else if len(oldSnapshotKeys) > 0 {
				for _, oldSnapshotkey := range oldSnapshotKeys {
					if err1 := ProcessSnapshot(oldSnapshotkey); err1 != nil {
						log.Println("There was some error in Processing the PACS002 Batch for Key", oldSnapshotkey, ", Error:", err1)
					}
				}
			} else {
				if snapshotKey, err := SnapshotBatch(); err != nil {
					if err.Error() != "ERR no such key" {
						log.Println("Failed to take a snapshot of Batch:", err)
					}
				} else {
					if err1 := ProcessSnapshot(snapshotKey); err1 != nil {
						log.Println("There was some error in Processing the PACS002 Batch for Key", snapshotKey, ", Error:", err1)
					} else {
						log.Println(snapshotKey, "processed.")
					}
				}
			}
		}
	}()
}

func AddPoToPacs002Batch(po string) error {
	_, err := cache.RedisClient.RPush(ctx, "PACS002BATCH", po).Result()
	if err != nil {
		log.Println("Failed to add PO to PACS002 batch:", err)
		return fmt.Errorf("Failed to add PO to PACS002 batch: %w", err)
	}
	return nil
}

func SnapshotBatch() (string, error) {
	snapshotName := "PACS002BATCH_" + time.DateTime
	if err := cache.RedisClient.Rename(ctx, "PACS002BATCH", snapshotName).Err(); err != nil {
		if strings.Contains(err.Error(), "ERR no such key") {
			return "", fmt.Errorf("ERR no such key")
		}
		log.Println("Failed to rename batch:", err)
		return "", fmt.Errorf("Failed to rename batch: %w", err)
	} else {
		return snapshotName, nil
	}
}

func hasOldSnapshots() ([]string, error) {
	if keys, err := cache.RedisClient.Keys(ctx, "PACS002BATCH_*").Result(); err != nil {
		log.Println("Failed to fetch snapshot keys:", err)
		return nil, fmt.Errorf("Failed to fetch snapshot keys: %w", err)
	} else {
		if len(keys) > 0 {
			log.Println("Found Old PACS002 snapshots pending for processing:", keys)
		}
		return keys, nil
	}
}

func ProcessSnapshot(snapshotKey string) error {
	pos, err := cache.RedisClient.LRange(ctx, snapshotKey, 0, -1).Result()
	if err != nil {
		log.Println("Failed to read snapshot:", err)
		return fmt.Errorf("Failed to read snapshot: %w", err)
	}
	fmt.Println("Processing Snapshot for POs:", pos)
	if cursor, err := db.DB.Collection("PaymentOrders").Find(ctx, bson.M{"entityid": bson.M{"$in": pos}}); err != nil {
		log.Println("Failed to fetch Cursor for", pos, ":", err)
		return fmt.Errorf("Failed to fetch Cursor for %v: %w", pos, err)
	} else {
		var poData []*models.PaymentOrder
		if err1 := cursor.All(ctx, &poData); err1 != nil {
			log.Println("Failed to fetch Payment Orders for", pos, ":", err)
			return fmt.Errorf("Failed to fetch Payment Orders for %v: %w", pos, err)
		}
		statusMap := make(map[string]string)
		for _, po := range poData {
			statusMap[po.TransactionId] = po.Status
		}

		pacs002Id, _ := gonanoid.New(12)
		pacs002 := mapping.MapToPacs002(poData, statusMap)
		pacs002.FIToFIPmtStsRpt.GrpHdr.MsgId = pacs002Id
		pacs002.FIToFIPmtStsRpt.GrpHdr.CreDtTm = time.Now().Format("2026-01-29T06:16:00Z")
		pacs002.Xmlns = "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.15"
		pacs002.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
		pacs002.SchemaLocation = pacs002.Xmlns + " schema.xsd"

		if xmlBytes, err := xml.Marshal(pacs002); err != nil {
			log.Println("Failed to generate PACS002 XML, Error:", err)
			return fmt.Errorf("Failed to generate PACS002 XML, Error: %w", err)
		} else {
			xmlWithHeader := []byte(xml.Header + string(xmlBytes))
			if err := routing.SendPacs002Out(string(xmlWithHeader), Pacs002CallbackUrl); err != nil {
				log.Println("Error in PACS002 Posting:", err)
				return fmt.Errorf("Error in PACS002 Posting: %w", err)
			} else {
				log.Println("PACS002 for all POs in Snapshot", snapshotKey, " Sent out.")
				deletedKeys, _ := cache.RedisClient.Del(ctx, snapshotKey).Result()
				log.Println(snapshotKey, "deleted:", deletedKeys)
				if res, err := db.DB.Collection("MessageLogger").InsertOne(ctx, map[string]any{
					"_id":        pacs002Id,
					"actualType": strings.Split(pacs002.Xmlns, "xsd:")[1],
					"route":      "outgoing",
					"asIsMsg":    string(xmlWithHeader),
					"createdAt":  time.Now(),
				}); err != nil {
					log.Println("There was some error in creating MessageLogger entry for PACS002 Out:", pacs002Id)
				} else {
					log.Println("MessageLogger entry for PACS002 Out created:", res)
				}
				return nil
			}
		}
	}
}

func CreatePacs002ForSinglePo(po *models.PaymentOrder, status string) error {
	pacs002Id, _ := gonanoid.New(12)
	pacs002 := mapping.MapToPacs002(append([]*models.PaymentOrder{}, po), map[string]string{po.TransactionId: status})
	pacs002.FIToFIPmtStsRpt.GrpHdr.MsgId = pacs002Id
	pacs002.FIToFIPmtStsRpt.GrpHdr.CreDtTm = time.Now().Format("2026-01-29T06:16:00Z")
	pacs002.Xmlns = "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.15"
	pacs002.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	pacs002.SchemaLocation = pacs002.Xmlns + " schema.xsd"

	if xmlBytes, err := xml.Marshal(pacs002); err != nil {
		log.Println("Failed to generate PACS002 XML, Error:", err)
		return fmt.Errorf("Failed to generate PACS002 XML, Error: %w", err)
	} else {
		xmlWithHeader := []byte(xml.Header + string(xmlBytes))
		if err := routing.SendPacs002Out(string(xmlWithHeader), Pacs002CallbackUrl); err != nil {
			log.Println("Error in PACS002 Posting:", err, " for PoId:", po.Id)
			return fmt.Errorf("Error in PACS002 Posting: %w for PoId: %v", err, po.Id)
		} else {
			log.Println("PACS002", pacs002Id, "for TransactionId:", po.Id, "Sent.")
			if res, err := db.DB.Collection("MessageLogger").InsertOne(ctx, map[string]any{
				"_id":        pacs002Id,
				"actualType": strings.Split(pacs002.Xmlns, "xsd:")[1],
				"route":      "outgoing",
				"asIsMsg":    string(xmlWithHeader),
				"createdAt":  time.Now(),
			}); err != nil {
				log.Println("There was some error in creating MessageLogger entry for PACS002 Out:", pacs002Id)
			} else {
				log.Println("MessageLogger entry for PACS002 Out created:", res)
			}
			return nil
		}
	}
}
