package web

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/shaninalex/homefilestorage/web/templates"
)

func (web *WebApp) homeHandler(w http.ResponseWriter, r *http.Request) {
	web.State.CSRFToken = csrf.Token(r)
	templates.Home(*web.State).Render(r.Context(), w)
}

func (web *WebApp) loginHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Actual user login
	web.State.LoggedIn = true
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (web *WebApp) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Actual user logout ( remove user object from state )
	web.State.LoggedIn = false
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
