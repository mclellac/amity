package api

import "time"

type Post struct {
	Id        int32     `gorm:"primary_key;column:id" json:"id"`
	Created   int32     `sql:"size:25"                  json:"created,omitempty"`
	Title     string    `sql:"size:140;not null"        json:"title" binding:"required"`
	Article   string    `sql:"size:8000;not null"       json:"article"  binding:"required"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}
