package processing

import (
	"encoding/json"
	"log"
	"payment-simulator/internal/models"
)

func ProcessInboundPo(po *models.PaymentOrder) {
	jsonData, _ := json.MarshalIndent(&po, "", " ")
	log.Println("Payment Order Rcvd: ", string(jsonData))
	if res, err := DuplicateCheck(po); res {
		po.Errors = append(po.Errors, err.Error())
		po.EntityId = "DUPLICATE"
		po.Status = "RJCT"
	} else {
		if err != nil {
			log.Printf("There was some error in Checking if PO %s is Duplicate: ", po.EntityId)
		}
		po.EntityId = AssignPoNumber()
		po.Status = "ACTC"
	}
}
