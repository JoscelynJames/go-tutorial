package models 

import (
	"github.com/inzhu/gorm"
)

type Book struct {
	ID uint `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Author string `json: author`
}
