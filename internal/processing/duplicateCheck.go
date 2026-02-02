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
	log.Println("Checking if PO is duplicate: ", po.PoNumber)
	dupCheckStr := strings.ToLower(po.Currency + fmt.Sprintf("%g", po.Amount) + po.Rcvr + po.RcvrAcct + po.Sender + po.SenderAcct)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if dupCheckFlag, err := cache.RedisClient.Exists(ctx, dupCheckStr).Result(); err != nil {
		log.Println("Error in Checking if PO is Duplicate:", err)
		return false, err
	} else if dupCheckFlag > 0 {
		errStr := fmt.Sprintf("PO %s is Duplicate.", po.PoNumber)
		log.Println(errStr)
		return true, errors.New(errStr)
	} else {
		log.Printf("\nPO %s is New.", po.PoNumber)
		if _, err := SetRecordForDupCheck(dupCheckStr, po.PoNumber); err != nil {
			return false, err
		} else {
			return false, nil
		}
	}
}
