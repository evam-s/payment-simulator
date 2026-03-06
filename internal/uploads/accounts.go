package uploads

import (
	"context"
	"encoding/json"
	"log"
	"payment-simulator/internal/db"
	"payment-simulator/internal/kafka"
	"payment-simulator/internal/reference"
	"time"
)

var ctxBg = context.Background()

func ConsumeFromAccountsTopic() {
	accountsChan := make(chan kafka.KafkaMsg)
	go kafka.ConsumeFromTopic("accounts", accountsChan)
	for accountMsg := range accountsChan {
		log.Println("msg from accounts Q", string(accountMsg.Value))
		var account reference.Account
		json.Unmarshal(accountMsg.Value, &account)
		log.Println("json from Q", account)

		ctx, cancel := context.WithTimeout(ctxBg, 5*time.Second)
		if res, err := db.DB.Collection("Accounts").InsertOne(ctx, account); err != nil {
			log.Println("There was some error in Creating Account:", account, ", Error:", err)
		} else {
			log.Println("New Account entry created:", res)
		}
		cancel()
	}
}

// Id,Name,PhoneNumber,Email,Address,Balance,Currency,Active
