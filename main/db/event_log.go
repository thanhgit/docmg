package db

import "time"

type Event_log struct {
	ID        		int `gorm:"primary_key"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		*time.Time `sql:"index"`
	Document		Document	`gorm:"foreignkey:DocumentRefer"`
	DocumentRefer 	uint
	Author			User		`gorm:"foreignkey:AuthorRefer"`
	AuthorRefer		uint
	Event			string
}

func NewEvent_log(document Document, author User, event string) *Event_log  {
	return &Event_log{
		Document: document,
		Author:   author,
		Event:       event,
	}
}

func CreateEvent_log(event_log Event_log) {
	if GetInstance().NewRecord(event_log) {
		GetInstance().Create(&event_log)
		if GetInstance().NewRecord(event_log) {
			println("Fail")
		} else {
			println("Success")
		}
	}
}