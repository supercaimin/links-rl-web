package propertys

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/utils"
	"github.com/qor/render"
	"github.com/qor/responder"
)

var decoder = schema.NewDecoder()

// Controller products controller
type Controller struct {
	View *render.Render
}

// Buy page
func (ctrl Controller) Buy(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("buy", map[string]interface{}{}, req, w)
}

// Rent page
func (ctrl Controller) Rent(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("rent", map[string]interface{}{}, req, w)
}

// Join page
func (ctrl Controller) Join(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("join", map[string]interface{}{}, req, w)
}

// About page
func (ctrl Controller) About(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("about", map[string]interface{}{}, req, w)
}

// Contact page
func (ctrl Controller) Contact(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("contact", map[string]interface{}{}, req, w)
}

// Show page
func (ctrl Controller) Show(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("show", map[string]interface{}{}, req, w)
}

// Sumbmit property
func (ctrl Controller) SumbmitProperty(w http.ResponseWriter, req *http.Request) {
	var (
		input propertys.ListProperty
		tx    = utils.GetDB(req)
	)

	req.ParseForm()
	decoder.Decode(&input, req.PostForm)

	tx.NewRecord(input)

	responder.With("html", func() {
		http.Redirect(w, req, "/", http.StatusFound)
	}).With([]string{"json", "xml"}, func() {
		config.Render.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}).Respond(req)
}
