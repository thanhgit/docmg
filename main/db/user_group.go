package db

import "time"

type User_group struct {
	ID         int `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	User       User	`gorm:"foreignkey:UserRefer"`
	UserRefer  uint
	Group      GroupUser `gorm:"foreignkey:GroupRefer"`
	GroupRefer uint
}

func NewUser_group(user User, group GroupUser) *User_group  {
	return &User_group{
		User:      user,
		Group:  group,
	}
}

func CreateUser_group(user_group User_group) {
	if GetInstance().NewRecord(user_group) {
		GetInstance().Create(&user_group)
		if GetInstance().NewRecord(user_group) {
			println("Fail")
		} else {
			println("Success")
		}
	}
}