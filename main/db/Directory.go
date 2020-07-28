package db

import "time"

type Directory struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name			string
	Ancestors		[]*Directory `gorm:"foreignkey:ID"`
	Parent			*Directory	`gorm:"foreignkey:ParentRef"`
	ParentRef		int
}

func NewDirectory(_name string, _ancestors []*Directory, _parent *Directory) *Directory {
	return &Directory{
		Name:      _name,
		Ancestors: _ancestors,
		Parent:    _parent,
	}
}