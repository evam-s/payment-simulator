package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Transaction struct {
	Sender string  `json:"payer"`
	Rcvr   string  `json:"payee"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
	payer2 string
}

func main() {
	fmt.Println("Payments Simulator starting...")
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
		c.JSON(200, gin.H{"status": "ok2"})
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.POST("/transaction", func(c *gin.Context) {
		fmt.Println("Payment Order Rcvd. Context: ", c)
		var tx Transaction
		if err := c.ShouldBindJSON(&tx); err != nil {
			fmt.Println("ERRRRRRRRRRRRRRRRRRRRR: ", tx)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tx.Status = "wait wat?"
		tx.payer2 = "asdfzxcv "
		fmt.Println("Payee: ", tx.Rcvr)
		fmt.Println("Payer: ", tx.Sender)
		fmt.Println("Amount: ", tx.Amount)
		fmt.Println("payer2: ", tx.payer2)
		c.JSON(200, gin.H{"message": "Transaction Rcvd", "transaction": tx})
	})
	router.Run(":8080")
}
