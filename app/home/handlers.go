package home

import (
	"net/http"

	"github.com/qor/qor"
	eutils "github.com/qor/qor-example/utils"
	"github.com/qor/qor/utils"
	"github.com/qor/render"
)

// Controller home controller
type Controller struct {
	View *render.Render
}

// Index home index page
func (ctrl Controller) Index(w http.ResponseWriter, req *http.Request) {
	var (
		Propertys1 []propertys.Property
		Propertys2 []propertys.Property
		Propertys3 []propertys.Property

		tx = eutils.GetDB(req)
	)

	tx.Limit(5).Find(&Propertys1, "is_premier_property=? AND is_vaild=?", 1, 1)
	tx.Offset(5).Limit(5).Find(&Propertys2, "is_premier_property=? AND is_vaild=?", 1, 1)

	tx.Order("view_count desc").Limit(6).Find(&Propertys3, "is_premier_property=? AND is_vaild=?", 0, 1)

	ctrl.View.Execute("index", map[string]interface{}{"Propertys1": Propertys1, "Propertys2": Propertys2, "Propertys3": Propertys3}, req, w)
}

// SwitchLocale switch locale
func (ctrl Controller) SwitchLocale(w http.ResponseWriter, req *http.Request) {
	utils.SetCookie(http.Cookie{Name: "locale", Value: req.URL.Query().Get("locale")}, &qor.Context{Request: req, Writer: w})
	http.Redirect(w, req, req.Referer(), http.StatusSeeOther)
}
