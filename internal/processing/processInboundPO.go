package processing

import (
	"encoding/json"
	"fmt"
	// "payment-simulator/internal/cache"
	"payment-simulator/internal/models"
)

func ProcessInboundPo(po *models.PaymentOrder) {
	jsonData, _ := json.MarshalIndent(&po, "", " ")
	fmt.Println("Payment Order Rcvd: ", string(jsonData))
	po.PoNumber = AssignPoNumber()
	po.Status = "ACCC"
}
