package handler

import (
	"net/http"

	"github.com/alwindoss/akademy"
	"github.com/justinas/nosurf"
)

type NoSurf struct {
	Cfg *akademy.Config
}

func (n NoSurf) NoSurfMW(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   n.Cfg.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
