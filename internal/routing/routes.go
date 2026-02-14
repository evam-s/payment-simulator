package routing

import (
	"bytes"
	"context"
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"payment-simulator/internal/db"
	// "payment-simulator/internal/iso20022/isomodels"
	"github.com/gin-gonic/gin"
	nanoid "github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/bson"
	"payment-simulator/internal/mapping"
	"payment-simulator/internal/processing"
	"time"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func RoutingSetup() *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.POST("/inboundPacs008", func(c *gin.Context) {
		fmt.Println("Payment Order Rcvd. Context: ", c)
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		id, _ := nanoid.New(12)
		contextBg := context.Background()
		ctx, cancel := context.WithTimeout(contextBg, 5*time.Second)
		defer cancel()

		if res, err := db.DB.Collection("MessageLogger").InsertOne(ctx, map[string]any{
			"_id":         id,
			"asIsMsg":     string(bodyBytes),
			"received_at": time.Now(),
		}); err != nil {
			log.Printf("Failed to log raw payload: %v\n", err)
		} else {
			log.Println("Logged Transaction Input, Id: ", res.InsertedID)
		}

		if isoPacs, err := mapping.MapXmlPacs008(c); err != nil {
			ctx1, cancel1 := context.WithTimeout(contextBg, 5*time.Second)
			defer cancel1()
			update := bson.M{"$set": bson.M{"error": err.Error()}}
			db.DB.Collection("MessageLogger").UpdateByID(ctx1, id, update)

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx1, cancel1 := context.WithTimeout(contextBg, 5*time.Second)
			defer cancel1()

			// log.Println("isoPacs", isoPacs)
			update := bson.M{"$set": bson.M{"transformedXml": isoPacs}}
			db.DB.Collection("MessageLogger").UpdateByID(ctx1, id, update)

			if err := processing.ProcessInboundPo(isoPacs); err != nil {
				c.JSON(400, gin.H{"ErrorMessage": err})
			} else {
				c.JSON(200, gin.H{"PO": *isoPacs})
			}
			// opts := options.FindOne().
			// 	SetSort(bson.D{{Key: "received_at", Value: -1}}).
			// 	SetProjection(bson.M{"transformedXml": 1})

			// var wrapper struct {
			// 	TransformedXml isomodels.Pacs008 `bson:"transformedXml"`
			// }
			// err := db.DB.Collection("MessageLogger").
			// 	FindOne(contextBg, bson.M{}, opts).
			// 	Decode(&wrapper)

			// log.Println("err:", err)
			// by, _ := json.MarshalIndent(&wrapper, "", "  ")
			// log.Println("result:", string(by))

			// ctx1, cancel1 := context.WithTimeout(contextBg, 5*time.Second)
			// defer cancel1()
			// update := bson.M{"$set": po}
			// db.DB.Collection("MessageLogger").UpdateByID(ctx1, id, update)
			// log.Println("Processed Inbound Payment Order: ", *po)
		}
	})
	return router
}
