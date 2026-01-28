package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"payment-simulator/internal/config"
	"payment-simulator/internal/db"
	"payment-simulator/internal/filestore"
)

type Transaction struct {
	Sender     string  `xml:"Payer"`
	SenderAcct string  `xml:"PayerAcct"`
	Rcvr       string  `xml:"Payee"`
	RcvrAcct   string  `xml:"PayeeAcct"`
	Amount     float64 `xml:"Amount"`
	Status     string  `xml:"Status"`
}

func main() {
	config := config.LoadConfig()
	fmt.Println("Starting App in ", config.ServiceMode, " Mode...")
	db.ConnectMongo(config.DBTECH + "://" + config.DBURL + ":" + config.DBPORT)
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
		if err := c.ShouldBindXML(&tx); err != nil {
			fmt.Println("Error in Binding to XML for Transaction: ", tx)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Payee: ", tx.Rcvr)
		fmt.Println("Payer: ", tx.Sender)
		fmt.Println("Amount: ", tx.Amount)

		store := filestore.NewFileStore("bucket/data.json")
		if err := ensureDir("bucket/data.json"); err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
		txnId := (tx.Sender + "_" + tx.Rcvr + "_" + fmt.Sprint(rand.Intn(500)))
		tx.Status = "Payment Rcvd!, Id:" + txnId
		payments := []filestore.Payment{{Id: txnId, Amount: tx.Amount}}
		log.Printf("payments = %+v", payments)
		if err := store.Save(payments); err != nil {
			log.Printf("Save error: %v", err)
		}
		loaded, _ := store.Load()
		fmt.Println("\nMost Recent payment:", loaded[0])

		c.JSON(200, gin.H{"message": "Transaction Rcvd", "transaction": tx})
	})
	fmt.Println("Payments App Started on Port: ", config.ServicePort)
	router.Run(":" + config.ServicePort)
}

func ensureDir(path string) error {
	dir := filepath.Dir(path)
	return os.MkdirAll(dir, os.ModePerm)
}
