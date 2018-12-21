package propertys

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	Name        string
	CompanyName string
	Phone       string

	Email    string
	Content  string `sql:"type:text"`
	IsReview bool
}
