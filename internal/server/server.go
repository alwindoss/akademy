package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/alwindoss/akademy"
	"github.com/alwindoss/akademy/internal/handler"
	"github.com/gorilla/mux"
)

func Run(cfg *akademy.Config) error {
	tc, err := handler.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	cfg.TemplateCache = tc
	session := scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = cfg.InProduction

	router := mux.NewRouter()
	n := handler.NoSurf{
		Cfg: cfg,
	}
	router.Use(n.NoSurfMW)
	router.Use(session.LoadAndSave)

	hdlrs := handler.NewPageHandler(cfg, session, nil)
	setupRoutes(router, hdlrs)
	addr := fmt.Sprintf(":%d", cfg.Port)
	fmt.Printf("Listening on %s\n", addr)
	return http.ListenAndServe(addr, router)
}

func setupRoutes(r *mux.Router, h handler.PageHandler) {
	r.Path("/").Methods(http.MethodGet).HandlerFunc(h.ShowIndexPage)
	r.Path("/about").Methods(http.MethodGet).HandlerFunc(h.ShowAboutPage)
	r.Path("/login").Methods(http.MethodGet).HandlerFunc(h.ShowLoginPage)
}
