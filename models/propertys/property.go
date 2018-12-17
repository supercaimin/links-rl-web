package propertys

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/location"
	"github.com/qor/media/media_library"
	"github.com/qor/qor-example/models/users"
	"github.com/qor/sorting"
)

type District struct {
	gorm.Model
	l10n.Locale
	Name  string
	Areas Areas `sql:"type:text"`
}

type PropertyLayouts struct {
	gorm.Model
	l10n.Locale
	Name string
}
type BuildingType struct {
	gorm.Model
	l10n.Locale
	Name string
}
type FloorZone struct {
	gorm.Model
	l10n.Locale
	Name string
}

type NoOfBedRooms struct {
	gorm.Model
	l10n.Locale
	Name string
}

type NoOfBathRooms struct {
	gorm.Model
	l10n.Locale
	Name string
}

type PropertyViews struct {
	gorm.Model
	l10n.Locale
	Name string
}

type Condition struct {
	gorm.Model
	l10n.Locale
	Name string
}

type Facility struct {
	gorm.Model
	l10n.Locale
	Name string
}

type PropertyType struct {
	gorm.Model
	l10n.Locale
	Name string
}
type Outdoor struct {
	gorm.Model
	l10n.Locale
	Name string
}

type Room struct {
	gorm.Model
	l10n.Locale
	Name string
}

type Direction struct {
	gorm.Model
	l10n.Locale
	Name string
}

type Property struct {
	gorm.Model
	l10n.Locale
	sorting.SortingDESC

	MainImage media_library.MediaBox
	Images    media_library.MediaBox

	Owner            string
	ContactPersons   []users.User `l10n:"sync" gorm:"many2many:contact_persons;"`
	AgentReferenceNo string
	Remark           string
	SellingPrice     float64
	SellingPriceRank string
	AskingRent       float64
	Inclusive        bool
	ManagementFee    float64
	GovRates         float64

	location.Location `location:"name:Property Address"`
	PropertyType      PropertyType
	BuildingName      string
	Floor             string
	Unit              string
	PropertyLayouts   PropertyLayouts
	BuildingType      BuildingType
	FloorZone         FloorZone
	NoOfBedRooms      NoOfBedRooms
	FloorSpace        string
	SaleableArea      string
	OutdoorAread      string
	NoOfBathRooms     NoOfBathRooms
	PropertyViews     []PropertyViews
	Condition         []Condition

	Facitlities []Facility
	Outdoor     []Outdoor
	Rooms       []Room
	Direction   []Direction
}

type Areas []Area
type Area struct {
	Name  string
	IsHot bool
}

func (areas *Areas) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, areas)
	case string:
		if v != "" {
			return areas.Scan([]byte(v))
		}
	default:
		return errors.New("not supported")
	}
	return nil
}

func (areas Areas) Value() (driver.Value, error) {
	if len(areas) == 0 {
		return nil, nil
	}
	return json.Marshal(areas)
}
