package handler

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/alwindoss/wys"
)

type PageHandler interface {
	ShowIndexPage(w http.ResponseWriter, r *http.Request)
	ShowAboutPage(w http.ResponseWriter, r *http.Request)
	ShowLoginPage(w http.ResponseWriter, r *http.Request)
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
