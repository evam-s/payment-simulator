package processing

import (
	"context"
	"errors"
	"fmt"
	"log"
	"payment-simulator/internal/cache"
	"payment-simulator/internal/models"
	"strings"
	"time"
)

func DuplicateCheck(po *models.PaymentOrder) (bool, error) {
	log.Println("Checking if PO is duplicate.")
	dupCheckStr := strings.ToLower(po.SettlementAmount.Currency + fmt.Sprintf("%g", po.SettlementAmount.Value) + po.Creditor.Name + po.CreditorAcct.Iban + po.Debtor.Name + po.DebtorAcct.Iban)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if dupCheckFlag, err := cache.RedisClient.Exists(ctx, dupCheckStr).Result(); err != nil {
		log.Println("Error in Checking if PO is Duplicate:", err)
		return false, err
	} else if dupCheckFlag > 0 {
		errStr := "PO is Duplicate."
		log.Println(errStr)
		return true, errors.New(errStr)
	} else {
		log.Println("PO is New.")
		if _, err := SetRecordForDupCheck(dupCheckStr, po.EntityId); err != nil {
			return false, err
		} else {
			return false, nil
		}
	}
}
