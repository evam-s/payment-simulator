package processing

import (
	"fmt"
	"log"
	"net/http"
	"payment-simulator/internal/iso20022/isomodels"
	"payment-simulator/internal/mapping"
	"time"
)

func ProcessInboundPo(isoPacs *isomodels.Pacs008, id string, headers http.Header) error {
	if err, po := mapping.MapPacs008ToPo(isoPacs); err != nil {
		log.Println("There was some error in mapping PO", po.EntityId, "to internal structures:", err)
		return fmt.Errorf("Failed to bind XML: %w", err)
	} else {
		po.Id = id
		if res, err := DuplicateCheck(po); res {
			po.Errors = append(po.Errors, err.Error())
			po.EntityId = "DUPLICATE"
			po.Status = "RJCT"
		} else {
			if err != nil {
				log.Println("There was some error in Checking if PO", po.EntityId, "is Duplicate:", err)
				return err
			}

			po.Status = "ACTC"
			po.TxnAcceptanceDateTime = time.Now().Format("2006-01-02T15:04:05Z")
			createPaymentOrder(po)
			if err := AddPoToPacs002Batch(po.EntityId); err != nil {
				log.Println("There was some error in adding PO", po.EntityId, " to PACS002 Batch:", err)
				return err
			}
		}
		return nil
	}
}
