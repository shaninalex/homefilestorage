package web

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/shaninalex/homefilestorage/pkg/config"
	"github.com/shaninalex/homefilestorage/pkg/database"
	"github.com/shaninalex/homefilestorage/web/templates"
)

type WebApp struct {
	Context  context.Context
	Config   *config.Config
	Database database.Repository
	Router   *mux.Router
	State    *templates.State
	Store    *sessions.CookieStore
}

func InitializeWebApp(conf *config.Config, db database.Repository) (*WebApp, error) {
	webapp := &WebApp{
		Context:  context.TODO(),
		Config:   conf,
		Database: db,
		Router:   mux.NewRouter(),
		State:    &templates.State{LoggedIn: false},
		Store:    sessions.NewCookieStore([]byte(conf.Web.SecretKey)),
	}
	err := db.Migrate()
	if err != nil {
		return nil, err
	}
	go webapp.createAdmin()
	webapp.initializeRoutes()
	return webapp, nil
}

func (web *WebApp) Run(port int64) {
	csrfMiddleware := csrf.Protect(
		[]byte(web.Config.CSRF.CsrfString),
		csrf.RequestHeader("Authenticity-Token"),
		csrf.FieldName("authenticity_token"),
		csrf.SameSite(csrf.SameSiteStrictMode),
	)
	web.Router.Use(csrfMiddleware)

	log.Printf("App started on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), web.Router)
}

func (web *WebApp) initializeRoutes() {
	web.Router.HandleFunc("/", web.homeHandler).Methods("GET")
	web.Router.HandleFunc("/login", web.loginHandler).Methods("POST")
	web.Router.HandleFunc("/logout", web.logoutHandler).Methods("GET")
}

func (web *WebApp) createAdmin() error {
	_, err := web.Database.GetAccountByEmail(web.Config.Admin.Email)
	if err == nil {
		// Account exists
		return nil
	}

	account := &database.Account{Name: web.Config.Admin.Name, Email: web.Config.Admin.Email}
	account.HashPassword(web.Config.Admin.Password)
	_, err = web.Database.CreateAccount(account.Email, account.Name, account.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}
