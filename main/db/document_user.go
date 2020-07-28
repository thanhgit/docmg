package db

import "time"

type Document_user struct {
	ID        		int `gorm:"primary_key"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		*time.Time `sql:"index"`
	Document		Document	`gorm:"foreignkey:DocumentRefer"`
	DocumentRefer 	uint
	User			User		`gorm:"foreignkey:UserRefer"`
	UserRefer		uint
	Is_read			bool
	Is_written		bool
}

func NewDocument_user(document Document, user User, is_read bool, is_written bool) *Document_user  {
	return &Document_user{
		Document: 	document,
		User:     		user,
		Is_read:     	is_read,
		Is_written:  	is_written,
	}
}

func CreateDocument_user(doc_user Document_user) {
	if GetInstance().NewRecord(doc_user) {
		GetInstance().Create(&doc_user)
		if GetInstance().NewRecord(doc_user) {
			println("Fail")
		} else {
			println("Success")
		}
	}
}
