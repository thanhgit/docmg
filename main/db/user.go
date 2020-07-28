package db

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name string
}

func NewUser(name string)	*User  {
	return &User{
		Name:      name,
	}
}

func CreateUser(user User) {
	if GetInstance().NewRecord(user) {
		GetInstance().Create(&user)
		if GetInstance().NewRecord(user) {
			println("Fail")
		} else {
			println("Success")
		}
	}
}