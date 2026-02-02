package processing

import (
	"log"
	"os"
	"payment-simulator/internal/cache"
	"strconv"
	"time"
)

var duplicateTxnTtl time.Duration = 0

func init() {
	ttlStr := os.Getenv("PO_DUPLICATE_CHECK_TTL")

	if ttlStr != "" {
		if secs, err := strconv.Atoi(ttlStr); err == nil {
			duplicateTxnTtl = time.Duration(secs) * time.Second
		}
	}
	if duplicateTxnTtl == 0 {
		duplicateTxnTtl = time.Duration(1800) * time.Second
	}
}

func AssignPoNumber() string {
	return cache.GetNewPoNumber()
}

func SetRecordForDupCheck(record, poNumber string) (bool, error) {
	if res, err := cache.StoreUsingSetWithTtl(record, poNumber, duplicateTxnTtl); err != nil {
		log.Println("Error in Setting record for Duplicate Check:", err)
		return false, err
	} else {
		log.Println("Record for Duplicate Check Set:", res)
		return res, nil
	}
}
