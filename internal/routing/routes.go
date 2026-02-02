package routing

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	nanoid "github.com/matoous/go-nanoid/v2"
	"io"
	"log"
	"net/http"
	"payment-simulator/internal/db"
	"payment-simulator/internal/mapping"
	"payment-simulator/internal/processing"
	"time"
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
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if res, err := db.DB.Collection("transactionInput").InsertOne(ctx, map[string]any{
			"_id":         id,
			"raw_payload": string(bodyBytes),
			"received_at": time.Now(),
		}); err != nil {
			log.Printf("Failed to log raw payload: %v\n", err)
		} else {
			log.Println("Logged Transaction Input, Id: ", res.InsertedID)
		}

		if po, err := mapping.MapIsoPacs008ToPo(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			processing.ProcessInboundPo(po)
			log.Println("Process Inbound Payment Order: ", *po)
			c.JSON(200, gin.H{"PO": *po})
		}
	})
	return router
}
