package processing

import (
	"fmt"
	"log"
	"payment-simulator/internal/iso20022/isomodels"
	"payment-simulator/internal/mapping"
	"payment-simulator/internal/models"
	"time"
)

var Pacs002CallbackUrl string // no point in setting after rcvng a pacs008. on restart this will be empty till we get a new one

func SetPacs002CB(url string) {
	Pacs002CallbackUrl = url
}

func ProcessInboundPo(isoPacs *isomodels.Pacs008, id, callBackUrl string) error {
	Pacs002CallbackUrl = callBackUrl
	if err, po := mapping.MapPacs008ToPo(isoPacs); err != nil {
		log.Println("There was some error in mapping PO", po.EntityId, "to internal structures:", err)
		return fmt.Errorf("Failed to bind XML: %w", err)
	} else {
		po.Id = id

		CreateEventLog(models.EventLog{
			ReqId:       po.Id,
			EventId:     "INTMAPCOMP",
			EventRemark: "Internal Mapping Complete",
		})

		po.Status = "RCVD"
		if err1 := CreatePacs002ForSinglePo(po, "RCVD"); err1 != nil {
			log.Println("There was some error in Posting RCVD PACS002 for PO", po.EntityId, ", Error:", err1)
			return fmt.Errorf("There was some error in Posting RCVD PACS002 for TxId %v, Error: %w", po.TransactionId, err1)
		}

		CreateEventLog(models.EventLog{
			ReqId:       po.Id,
			EventId:     "POVALINIT",
			EventRemark: "Payment Order Validation Started",
		})

		if res, dupPoId, err := DuplicateCheck(po); res {
			po.Errors = append(po.Errors, err.Error())
			po.EntityId = dupPoId + "-DUPLICATE"
			po.Status = "RJCT"

			CreateEventLog(models.EventLog{
				ReqId:       po.Id,
				EventId:     "POVALCOMP",
				EventRemark: "Payment Order is Duplicate of PO with TxId " + dupPoId + ". Rejected",
			})

			if err1 := CreatePacs002ForSinglePo(po, "RJCT"); err1 != nil {
				log.Println("There was some error in Posting RJCT PACS002 for PO", po.EntityId, ", Error:", err1)
				return fmt.Errorf("There was some error in Posting RJCT PACS002 for TxId %v, Error: %w", po.TransactionId, err1)
			}

			CreateEventLog(models.EventLog{
				ReqId:       po.Id,
				EventId:     "PACS002SENT",
				EventRemark: "PACS002 RJCT is sent",
			})
		} else {
			if err != nil {
				log.Println("There was some error in Checking if PO", po.EntityId, "is Duplicate:", err)
				return err
			}

			po.Status = "ACTC"
			po.TxnAcceptanceDateTime = time.Now().Format("2006-01-02T15:04:05Z")

			CreateEventLog(models.EventLog{
				ReqId:       po.Id,
				EventId:     "POVALCOMP",
				EventRemark: "Payment Order Validation Completed",
			})

			createPaymentOrder(po)

			log.Println("Payment Order Created:", po.EntityId)

			CreateEventLog(models.EventLog{
				ReqId:       po.Id,
				EntityId:    po.EntityId,
				EventId:     "POCREATE",
				EventRemark: "Payment Order was Created Succesfully",
			})

			if err := AddPoToPacs002Batch(po.EntityId); err != nil {
				log.Println("There was some error in adding PO", po.EntityId, " to PACS002 Batch:", err)
				return err
			}

			go func() {
				CreateEventLog(models.EventLog{
					ReqId:       po.Id,
					EntityId:    po.EntityId,
					EventId:     "CPSUCCESS",
					EventRemark: "Customer Processing completed",
				})

				po.Status = "ACSP"
				updatePaymentOrder(po)

				if err := AddPoToPacs002Batch(po.EntityId); err != nil {
					log.Println("There was some error in adding PO", po.EntityId, " to PACS002 Batch:", err)
					// return err
				}

				po.Status = "ACCC"
				updatePaymentOrder(po)

				if err := AddPoToPacs002Batch(po.EntityId); err != nil {
					log.Println("There was some error in adding PO", po.EntityId, " to PACS002 Batch:", err)
					// return err
				}
				TriggerPacs002Batch()
			}()
		}
		return nil
	}
}
