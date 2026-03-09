package processing

import (
	"errors"
	"fmt"
	"log"
	"payment-simulator/internal/cache"
	"payment-simulator/internal/models"
	"strings"
)

func DuplicateCheck(po *models.PaymentOrder) (bool, string, error) {
	log.Println("Checking if PO is duplicate.")
	dupCheckStr := strings.ToLower(po.SettlementAmount.Currency + fmt.Sprintf("%g", po.SettlementAmount.Value) + po.Creditor.Name + po.CreditorAcct.Iban + po.Debtor.Name + po.DebtorAcct.Iban)
	if dupCheckFlag, err := cache.FetchUsingGet(dupCheckStr); err != nil && err.Error() != "redis: nil" {
		log.Println("Error in Checking if PO is Duplicate:", err)
		return false, "", err
	} else if dupCheckFlag != "" {
		errStr := "PO is Duplicate."
		log.Println(errStr)
		return true, dupCheckFlag, errors.New(errStr)
	} else {
		log.Println("PO is New.")
		if _, err := SetRecordForDupCheck(dupCheckStr, po.Id); err != nil {
			return false, "", err
		} else {
			return false, "", nil
		}
	}
}
