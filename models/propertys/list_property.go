package propertys

import (
	"github.com/jinzhu/gorm"
)

type ListProperty struct {
	gorm.Model
	PropertyName string
	Name         string
	Address      string
	Contact      string
	Details      string `sql:"type:text"`
	IsReview     bool
}
