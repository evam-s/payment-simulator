package models

import "time"

type EventLog struct {
	Id             string `bson:"_id"`
	ReqId          string
	EntityId       string
	EventDate      time.Time
	EventId        string
	EventRemark    string
	AdditionalInfo string
	Errors         []string
}
