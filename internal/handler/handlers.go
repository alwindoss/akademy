package handler

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/alwindoss/akademy"
)

type PageHandler interface {
	ShowIndexPage(w http.ResponseWriter, r *http.Request)
	ShowAboutPage(w http.ResponseWriter, r *http.Request)
}

func NewPageHandler(cfg *akademy.Config, sess *scs.SessionManager, svc interface{}) PageHandler {
	return pageHandler{
		Cfg:     cfg,
		SessMgr: sess,
	}
}

type pageHandler struct {
	Cfg     *akademy.Config
	SessMgr *scs.SessionManager
}

// ShowAboutPage implements PageHandler.
func (h pageHandler) ShowAboutPage(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	h.SessMgr.Put(r.Context(), "remote-ip", remoteIP)
	d := &TemplateData{
		Title: "Akademy | About",
	}
	renderTemplate(w, r, h.Cfg, "about.page.html", d)
}

// ShowIndexPage implements PageHandler.
func (h pageHandler) ShowIndexPage(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	h.SessMgr.Put(r.Context(), "remote-ip", remoteIP)
	d := &TemplateData{
		Title: "Akademy | Home",
	}
	renderTemplate(w, r, h.Cfg, "index.page.html", d)
}
