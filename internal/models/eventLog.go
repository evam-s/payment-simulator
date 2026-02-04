package models

type EventLog struct {
	EntityId    string
	EventId     string
	EventRemark string
	Errors      []string
	MsgId       string
}
