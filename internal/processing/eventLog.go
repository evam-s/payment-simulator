package processing

import (
	"context"
	"errors"
	"fmt"
	"log"
	"payment-simulator/internal/db"
	"payment-simulator/internal/models"
	"time"

	"github.com/matoous/go-nanoid/v2"
)

func CreateEventLog(event models.EventLog) error {
	if event.ReqId == "" {
		return errors.New("Event.ReqId must be present to create EventLog")
	} else {
		eventId, _ := gonanoid.New(12)
		event.Id = eventId
		event.EventDate = time.Now()
		insertionCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if res, err := db.DB.Collection("EventLog").InsertOne(insertionCtx, event); err != nil {
			log.Println("There was some error in Creating EventLog:", event, ", Error:", err)
			return fmt.Errorf("There was some error in Creating EventLog: %v , Error: %w", event, err)
		} else {
			log.Println("EventLog creation response:", res)
			return nil
		}
	}
}
