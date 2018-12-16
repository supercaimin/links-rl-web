package property

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/sorting"
)

type Tag struct {
	gorm.Model
	l10n.Locale
	sorting.Sorting
	Name    string
	TagType string
}
