package propertys

import (
	"github.com/qor/admin"
	"github.com/qor/qor-example/config/application"
	"github.com/qor/qor-example/models/propertys"
)

var PriceRanks = []string{"Thousands", "Millions"}
var Fors = []string{"Sell", "Rent"}

// New new home app
func New(config *Config) *App {
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	//controller := &Controller{View: render.New(&render.Config{AssetFileSystem: application.AssetFS.NameSpace("propertys")}, "app/propertys/views")}

	//funcmapmaker.AddFuncMapMaker(controller.View)
	app.ConfigureAdmin(application.Admin)

	//application.Router.Get("/products", controller.Index)
	//application.Router.Get("/products/{code}", controller.Show)
	//application.Router.Get("/{gender:^(men|women|kids)$}", controller.Gender)
	//application.Router.Get("/category/{code}", controller.Category)
}

// ConfigureAdmin configure admin interface
func (App) ConfigureAdmin(Admin *admin.Admin) {
	// Property Management
	Admin.AddMenu(&admin.Menu{Name: "Property Management", Priority: 1})
	Admin.AddMenu(&admin.Menu{Name: "Property Settings", Priority: 3})
	Admin.AddMenu(&admin.Menu{Name: "Messages", Priority: 2})
	Admin.AddMenu(&admin.Menu{Name: "Position Management", Priority: 4})

	Admin.AddResource(&propertys.Message{}, &admin.Config{Menu: []string{"Messages"}})
	Admin.AddResource(&propertys.ListProperty{}, &admin.Config{Menu: []string{"Messages"}})

	Admin.AddResource(&propertys.Banner{}, &admin.Config{Menu: []string{"Property Management"}})
	Admin.AddResource(&propertys.Company{}, &admin.Config{Name: "Company Setting", Menu: []string{"Site Management"}, Singleton: true, Priority: 1})

	position := Admin.AddResource(&propertys.Position{}, &admin.Config{Menu: []string{"Position Management"}})
	position.IndexAttrs("Title", "WageDemands", "WorkingAddress")
	position.Meta(&admin.Meta{Name: "Details", Config: &admin.RichEditorConfig{Plugins: []admin.RedactorPlugin{
		{Name: "medialibrary", Source: "/admin/assets/javascripts/qor_redactor_medialibrary.js"},
		{Name: "table", Source: "/vendors/redactor_table.js"},
	},
		Settings: map[string]interface{}{
			"medialibraryUrl": "/system/media_libraries",
		},
	}})

	// Add Property
	property := Admin.AddResource(&propertys.Property{}, &admin.Config{Menu: []string{"Property Management"}})
	//property.Meta(&admin.Meta{Name: "ContactPersons", Config: &admin.SelectManyConfig{SelectMode: "bottom_sheet"}})
	//property.Meta(&admin.Meta{Name: "ContactPersons", Config: &admin.SelectManyConfig{PrimaryField: "Name"}})
	//property.Meta(&admin.Meta{Name: "Owner", Label: "Username(Owner)"})
	//property.Meta(&admin.Meta{Name: "AgentReferenceNo", Label: "Agent's Reference No."})

	property.Meta(&admin.Meta{Name: "SellingPrice", Label: "Selling Price(HK$)"})
	property.Meta(&admin.Meta{Name: "SellingPriceRank", Label: "", Config: &admin.SelectOneConfig{Collection: PriceRanks, AllowBlank: false}})
	property.Meta(&admin.Meta{Name: "For", Label: "For", Config: &admin.SelectOneConfig{Collection: Fors, AllowBlank: false}})

	property.Meta(&admin.Meta{Name: "AskingRent", Label: "Asking Rent(HK$)"})
	property.Meta(&admin.Meta{Name: "Inclusive", Label: "Inclusive?"})
	property.Meta(&admin.Meta{Name: "ManagementFee", Label: "Management Fee(HK$)"})
	property.Meta(&admin.Meta{Name: "GovRates", Label: "Gov. Rates(HK$)"})

	property.Meta(&admin.Meta{Name: "PropertyType", Config: &admin.SelectOneConfig{AllowBlank: true}})
	property.Meta(&admin.Meta{Name: "PropertyLayouts", Config: &admin.SelectOneConfig{AllowBlank: true}})
	property.Meta(&admin.Meta{Name: "BuildingType", Label: "Building Type(Unit)", Config: &admin.SelectOneConfig{AllowBlank: true}})
	property.Meta(&admin.Meta{Name: "FloorZone", Config: &admin.SelectOneConfig{AllowBlank: true}})
	property.Meta(&admin.Meta{Name: "NoOfBedRooms", Label: "No. Of Bedrooms", Config: &admin.SelectOneConfig{AllowBlank: true}})
	property.Meta(&admin.Meta{Name: "NoOfBathRooms", Label: "No. Of Bathdrooms", Config: &admin.SelectOneConfig{AllowBlank: true}})

	//property.Meta(&admin.Meta{Name: "FloorSpace", Label: "Floor Space(sq.ft.)"})
	property.Meta(&admin.Meta{Name: "GrossArea", Label: "Gross Area(sq.ft.)"})
	property.Meta(&admin.Meta{Name: "SaleableArea", Label: "Saleable Area(sq.ft.)"})
	property.Meta(&admin.Meta{Name: "OutdoorArea", Label: "Outdoor Area(sq.ft.)"})

	property.Meta(&admin.Meta{Name: "PropertyViews", Type: "select_many"})
	//property.Meta(&admin.Meta{Name: "Condition", Type: "select_many"})
	property.Meta(&admin.Meta{Name: "Facitlities", Type: "select_many"})
	property.Meta(&admin.Meta{Name: "Outdoor", Type: "select_many"})
	property.Meta(&admin.Meta{Name: "Rooms", Type: "select_many"})
	property.Meta(&admin.Meta{Name: "Direction", Type: "select_many"})

	//property.UseTheme("grid")

	district := Admin.AddResource(&propertys.District{}, &admin.Config{Menu: []string{"Property Settings"}})

	areas := district.Meta(&admin.Meta{Name: "Areas", Label: "Regions"}).Resource
	areas.NewAttrs(&admin.Section{
		Rows: [][]string{{"Name", "IsHot"}},
	})
	areas.EditAttrs(&admin.Section{
		Rows: [][]string{{"Name", "IsHot"}},
	})
	district.IndexAttrs("Name")
	district.EditAttrs("Name", "Areas")
	district.NewAttrs(district.EditAttrs())

	//buildingType := Admin.AddResource(&propertys.BuildingType{}, &admin.Config{Menu: []string{"Property Settings"}})

	//buildingType.EditAttrs("Name")
	//buildingType.NewAttrs(buildingType.EditAttrs())

	direction := Admin.AddResource(&propertys.Direction{}, &admin.Config{Menu: []string{"Property Settings"}})

	direction.EditAttrs("Name")
	direction.NewAttrs(direction.EditAttrs())

	facility := Admin.AddResource(&propertys.Facility{}, &admin.Config{Menu: []string{"Property Settings"}})

	facility.EditAttrs("Name")
	facility.NewAttrs(facility.EditAttrs())

	//floorZone := Admin.AddResource(&propertys.FloorZone{}, &admin.Config{Menu: []string{"Property Settings"}})

	//floorZone.EditAttrs("Name")
	//floorZone.NewAttrs(floorZone.EditAttrs())

	noOfBathRooms := Admin.AddResource(&propertys.NoOfBathRooms{}, &admin.Config{Menu: []string{"Property Settings"}})

	noOfBathRooms.EditAttrs("Name")
	noOfBathRooms.NewAttrs(noOfBathRooms.EditAttrs())

	noOfBedRooms := Admin.AddResource(&propertys.NoOfBedRooms{}, &admin.Config{Menu: []string{"Property Settings"}})

	noOfBedRooms.EditAttrs("Name")
	noOfBedRooms.NewAttrs(noOfBedRooms.EditAttrs())

	outdoor := Admin.AddResource(&propertys.Outdoor{}, &admin.Config{Menu: []string{"Property Settings"}})

	outdoor.EditAttrs("Name")
	outdoor.NewAttrs(outdoor.EditAttrs())

	propertyLayouts := Admin.AddResource(&propertys.PropertyLayouts{}, &admin.Config{Menu: []string{"Property Settings"}})

	propertyLayouts.EditAttrs("Name")
	propertyLayouts.NewAttrs(propertyLayouts.EditAttrs())

	propertyViews := Admin.AddResource(&propertys.PropertyViews{}, &admin.Config{Menu: []string{"Property Settings"}})

	propertyViews.EditAttrs("Name")
	propertyViews.NewAttrs(propertyViews.EditAttrs())

	room := Admin.AddResource(&propertys.Room{}, &admin.Config{Menu: []string{"Property Settings"}})

	room.EditAttrs("Name")
	room.NewAttrs(room.EditAttrs())

	propertyType := Admin.AddResource(&propertys.PropertyType{}, &admin.Config{Menu: []string{"Property Settings"}})

	propertyType.EditAttrs("Name")
	propertyType.NewAttrs(propertyType.EditAttrs())

	property.IndexAttrs("MainImage", "For", "SellingPrice", "AskingRent", "Location")

	property.EditAttrs(
		// &admin.Section{
		// 	Title: "Property Information",
		// 	Rows: [][]string{
		// 		{"MainImage"},
		// 		{"Owner", "AgentReferenceNo"},
		// 		{"ContactPersons", "For"},
		// 		{"Remark"},
		// 	}},
		&admin.Section{
			Title: "Selling Price",
			Rows: [][]string{
				{"SellingPrice", "SellingPriceRank"},
				{"AskingRent", "Inclusive", "For"},
			//	{"ManagementFee", "GovRates"},
			}},
		&admin.Section{
			Title: "Location",
			Rows: [][]string{
				{"Location"},
				{"PropertyType"},
				{"BuildingName", "Floor", "Unit"},
			}},
		&admin.Section{
			Title: "Details",
			Rows: [][]string{
				{"PropertyLayouts", "NoOfBedRooms"},
				//{"FloorZone", "BuildingType"},
				{"GrossArea", "SaleableArea"},
				{"OutdoorArea", "NoOfBathRooms"},
				{"PropertyViews"},
			}},
		&admin.Section{
			Title: "Features",
			Rows: [][]string{
				{"Facitlities", "Outdoor"},
				{"Rooms"},
			}},
		&admin.Section{
			Title: "Upload Images",
			Rows: [][]string{
				{"MainImage"},
				{"Images"},
			}},
		&admin.Section{
			Title: "Others",
			Rows: [][]string{
				{"IsPremierProperty", "IsVaild"},
			}},
	)

	property.NewAttrs(property.EditAttrs())

}
