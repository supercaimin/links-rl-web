package propertys

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/location"
	"github.com/qor/sorting"
)

type Company struct {
	gorm.Model
	Phone string
	Email string
	location.Location
	sorting.Sorting
}
