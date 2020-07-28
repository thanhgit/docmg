package db

import (
	"time"
)

type File struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name			string
	Content			string
	Parent		Directory `gorm:"foreignkey:ParentRef"`
	ParentRef		int
}

func NewFile(_name string, _content string, _directory Directory) *File {
	return &File{
		Name:      _name,
		Content:   _content,
		Parent: _directory,
	}
}