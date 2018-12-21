package propertys

import (
	"github.com/jinzhu/gorm"
)

type ListProperty struct {
	gorm.Model
	Name     string
	Address  string
	Contact  string
	Details  string `sql:"type:text"`
	IsReview bool
}
