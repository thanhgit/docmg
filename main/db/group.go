package db

import "time"

type GroupUser struct {
	ID        		int `gorm:"primary_key"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		*time.Time `sql:"index"`
	Name 			string
}

func NewGroup(name string) *GroupUser {
	return &GroupUser{
		Name:      name,
	}
}

func CreateGroup(group GroupUser) {
	if GetInstance().NewRecord(group) {
		GetInstance().Create(&group)
		if GetInstance().NewRecord(group) {
			println("Fail")
		} else {
			println("Success")
		}
	}
}
