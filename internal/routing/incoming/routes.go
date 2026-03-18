package routing

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"payment-simulator/internal/db"
	"payment-simulator/internal/mapping"
	"payment-simulator/internal/processing"
	"payment-simulator/internal/routing/graphql"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func RoutingSetup() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.POST("/inboundPacs008", func(c *gin.Context) {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		headers := c.Request.Header

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}
		log.Println("Request Headers:", headers)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		id, _ := gonanoid.New(12)
		contextBg := context.Background()
		ctx, cancel := context.WithTimeout(contextBg, 5*time.Second)
		defer cancel()

		if res, err := db.DB.Collection("MessageLogger").InsertOne(ctx, map[string]any{
			"_id":          id,
			"expectedType": "pacs008",
			"route":        "incoming",
			"asIsMsg":      string(bodyBytes),
			"receivedAt":   time.Now(),
		}); err != nil {
			log.Println("Failed to log raw payload:", err)
		} else {
			log.Println("Logged Transaction Input, Id:", res.InsertedID)
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
			update := bson.M{"$set": bson.M{"transformedXml": isoPacs, "actualType": strings.Split(isoPacs.Xmlns, "xsd:")[1]}}
			db.DB.Collection("MessageLogger").UpdateByID(ctx1, id, update)

			if err := processing.ProcessInboundPo(isoPacs, id); err != nil {
				if strings.HasPrefix(err.Error(), "Failed to bind XML") {
					c.JSON(400, gin.H{"ErrorMessage": err})
				} else {
					c.JSON(503, gin.H{"ErrorMessage": err})
				}
			} else {
				// c.JSON(200, gin.H{"PO": *isoPacs})
				c.XML(200, *isoPacs)
			}
		}
	})
	graphql.GraphQlSetup(router)
	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
