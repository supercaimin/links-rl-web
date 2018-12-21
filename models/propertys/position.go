package propertys

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
)

type Position struct {
	gorm.Model
	l10n.Locale
	Title          string
	WageDemands    string
	WorkingAddress string
	Details        string `sql:"type:text"`
}
