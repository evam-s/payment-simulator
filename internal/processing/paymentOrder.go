package processing

import (
	"context"
	"log"
	"os"
	"payment-simulator/internal/cache"
	"payment-simulator/internal/db"
	"payment-simulator/internal/models"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var duplicateTxnTtl time.Duration = 0

func init() {
	ttlStr := os.Getenv("PO_DUPLICATE_CHECK_TTL")

	if ttlStr != "" {
		if secs, err := strconv.Atoi(ttlStr); err == nil {
			duplicateTxnTtl = time.Duration(secs) * time.Second
		}
	}
	if duplicateTxnTtl == 0 {
		duplicateTxnTtl = time.Duration(1800) * time.Second
	}
}

func AssignPoNumber() string {
	return cache.GetNewPoNumber()
}

func SetRecordForDupCheck(record, poNumber string) (bool, error) {
	if res, err := cache.StoreUsingSetWithTtl(record, poNumber, duplicateTxnTtl); err != nil {
		log.Println("Error in Setting record for Duplicate Check:", err)
		return false, err
	} else {
		log.Println("Record for Duplicate Check Set:", res)
		return res, nil
	}
}

func createPaymentOrder(po *models.PaymentOrder) error {
	po.EntityId = AssignPoNumber()
	po.CreatedOn = time.Now()
	contextBg := context.Background()
	ctx, cancel := context.WithTimeout(contextBg, 5*time.Second)
	defer cancel()
	if res, err := db.DB.Collection("PaymentOrders").InsertOne(ctx, po); err != nil {
		log.Println("Failed to Create Payment Order Entry:", err)
		return err
	} else {
		log.Println("Payment Order Entry Created. Id:", res.InsertedID)
		return nil
	}
}

func updatePaymentOrder(po *models.PaymentOrder) error {
	if res, err := db.DB.Collection("PaymentOrders").UpdateOne(ctx, bson.M{"entityid": po.EntityId}, bson.M{"$set": po}); true {
		log.Println("PaymentOrders UpdateOne res", res)
		log.Println("PaymentOrders UpdateOne err", err)
	}
	return nil
}
