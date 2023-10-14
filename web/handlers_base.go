package web

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/shaninalex/homefilestorage/web/templates"
)

func (web *WebApp) homeHandler(w http.ResponseWriter, r *http.Request) {
	web.State.CSRFToken = csrf.Token(r)
	session, _ := web.Store.Get(r, "session.id")
	authenticated := session.Values["authenticated"]
	if authenticated != nil && authenticated != false {
		web.State.LoggedIn = true
		web.State.ActiveRoute = "files"
		files, _ := web.Database.AllFiles()
		web.State.Files = files
	} else {
		web.State.LoggedIn = false
	}
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

	session, _ := web.Store.Get(r, "session.id")
	session.Values["authenticated"] = true
	session.Save(r, w)

	web.State.LoggedIn = true
	web.State.Error = ""
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (web *WebApp) logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := web.Store.Get(r, "session.id")
	// Set the authenticated value on the session to false
	session.Values["authenticated"] = false
	session.Save(r, w)
	web.State.LoggedIn = false
	templates.Home(*web.State).Render(r.Context(), w)
}
