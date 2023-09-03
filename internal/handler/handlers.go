package handler

import (
	"fmt"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/alwindoss/wys"
)

type PageHandler interface {
	// Renders pages
	ShowIndexPage(w http.ResponseWriter, r *http.Request)
	ShowAboutPage(w http.ResponseWriter, r *http.Request)
	ShowLoginPage(w http.ResponseWriter, r *http.Request)

	// Handles requests from the browser
	HandleLogin(w http.ResponseWriter, r *http.Request)
}

func NewPageHandler(v wys.ViewManager, sess *scs.SessionManager, svc interface{}) PageHandler {
	return pageHandler{
		SessMgr: sess,
		ViewMgr: v,
	}
}

type pageHandler struct {
	// Cfg     *akademy.Config
	SessMgr *scs.SessionManager
	ViewMgr wys.ViewManager
}

// HandleLogin implements PageHandler.
func (h pageHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Login")
	r.ParseForm()
	email := r.FormValue("emailAddress")
	fmt.Println("Email", email)
	w.Write([]byte(email))
}

// ShowLoginPage implements PageHandler.
func (h pageHandler) ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	d := &wys.TemplateData{
		Title: "Akademy | Login",
	}
	h.ViewMgr.Render(w, r, "login.page.html", d)
}

// ShowAboutPage implements PageHandler.
func (h pageHandler) ShowAboutPage(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	h.SessMgr.Put(r.Context(), "remote-ip", remoteIP)
	d := &wys.TemplateData{
		Title: "Akademy | About",
	}
	h.ViewMgr.Render(w, r, "about.page.html", d)
}

// ShowIndexPage implements PageHandler.
func (h pageHandler) ShowIndexPage(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	h.SessMgr.Put(r.Context(), "remote-ip", remoteIP)
	d := &wys.TemplateData{
		Title: "Akademy | Home",
	}
	h.ViewMgr.Render(w, r, "index.page.html", d)
}
