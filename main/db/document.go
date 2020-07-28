package db

import (
	"time"
)

type Document struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Author_id 	int
	Parent_id  	int
	Obj_type	string
	Title		string
	Level		int
}

func NewDocument(author_id int, parent_id int, obj_tyoe string, title string) *Document {
	return &Document{
		Author_id: author_id,
		Parent_id: parent_id,
		Obj_type:  obj_tyoe,
		Title:     title,
	}
}

func CreateDocument(doc Document) {
	if GetInstance().NewRecord(doc) {
		GetInstance().Create(&doc)
		if GetInstance().NewRecord(doc) {
			println("Fail")
		} else {
			println("Success")
		}
	}
}