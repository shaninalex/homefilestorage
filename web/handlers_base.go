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
	err := r.ParseForm()
	if err != nil {
		web.State.Error = "Unable to parse login form"
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	v := r.Form
	email := v.Get("email")
	password := v.Get("password")
	account, err := web.Database.GetAccountByEmail(email)
	if err != nil {
		web.State.Error = "User not found"
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	if !account.CheckPassword(password) {
		web.State.Error = "Password does not match"
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	web.State.LoggedIn = true
	web.State.Error = ""
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (web *WebApp) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Actual user logout ( remove user object from state )
	web.State.LoggedIn = false
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
