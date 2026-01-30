package routing

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-simulator/internal/mapping"
	"payment-simulator/internal/processing"
)

func RoutingSetup() *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.POST("/transaction", func(c *gin.Context) {
		fmt.Println("Payment Order Rcvd. Context: ", c)
		if po, err := mapping.MapIsoPacs008ToPo(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			processing.ProcessInboundPO(po)
			c.JSON(200, gin.H{"PO": &po})
		}
	})
	return router
}
