package processing

import (
	// "encoding/json"
	"log"
	"payment-simulator/internal/iso20022/isomodels"
	"payment-simulator/internal/mapping"
)

func ProcessInboundPo(isoPacs *isomodels.Pacs008) error {
	if err, _ := mapping.MapPacs008ToPo(isoPacs); err != nil {
		// if err, po := mapping.MapPacs008ToPo(isoPacs); err != nil {
		log.Println("mapping.MapPacs008ToPo err:", err)

		return err
	} else {
		// jsonData, _ := json.MarshalIndent(po, "", " ")
		// log.Println("Payment Order Rcvd: ", string(jsonData))

		// if res, err := DuplicateCheck(po); res {
		// 	po.Errors = append(po.Errors, err.Error())
		// 	po.EntityId = "DUPLICATE"
		// 	po.Status = "RJCT"
		// } else {
		// 	if err != nil {
		// 		log.Printf("There was some error in Checking if PO %s is Duplicate: ", po.EntityId)
		// 	}
		// 	po.EntityId = AssignPoNumber()
		// 	po.Status = "ACTC"
		// }
		return nil
	}
}
