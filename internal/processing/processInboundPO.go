package processing

import (
	"fmt"
	// "github.com/gin-gonic/gin"
	"payment-simulator/internal/models"
)

func ProcessInboundPO(po *models.PaymentOrder) *models.PaymentOrder {
	fmt.Println()
	// fmt.Println("Payee: ", tx.Rcvr)
	// fmt.Println("Payer: ", tx.Sender)
	// fmt.Println("Amount: ", tx.Amount)

	// store := filestore.NewFileStore("bucket/data.json")
	// if err := ensureDir("bucket/data.json"); err != nil {
	// 	log.Fatalf("Failed to create directory: %v", err)
	// }
	// txnId := (tx.Sender + "_" + tx.Rcvr + "_" + fmt.Sprint(rand.Intn(500)))
	// tx.Status = "Payment Rcvd!, Id:" + txnId
	// payments := []filestore.Payment{{Id: txnId, Amount: tx.Amount}}
	// log.Printf("payments = %+v", payments)
	// if err := store.Save(payments); err != nil {
	// 	log.Printf("Save error: %v", err)
	// }
	// loaded, _ := store.Load()
	// fmt.Println("\nMost Recent payment:", loaded[0])

	// c.JSON(200, gin.H{"message": "Transaction Rcvd", "transaction": tx})
	po.Status = "ACCC"
	return po
}
