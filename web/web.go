package web

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
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
}

func InitializeWebApp(conf *config.Config, db database.Repository) (*WebApp, error) {
	webapp := &WebApp{
		Context:  context.TODO(),
		Config:   conf,
		Database: db,
		Router:   mux.NewRouter(),
		State:    &templates.State{LoggedIn: false},
	}
	err := db.Migrate()
	if err != nil {
		return nil, err
	}
	webapp.initializeRoutes()
	return webapp, nil
}

func (web *WebApp) Run(port int64) {
	log.Printf("App started on port %d\n", port)
	// TODO: csrf token authKey config option
	csrfMiddleware := csrf.Protect(
		[]byte(web.Config.CSRF.CsrfString),
		csrf.RequestHeader("Authenticity-Token"),
		csrf.FieldName("authenticity_token"),
		csrf.SameSite(csrf.SameSiteStrictMode),
	)
	web.Router.Use(csrfMiddleware)
	http.ListenAndServe(fmt.Sprintf(":%d", port), web.Router)
}

func (web *WebApp) initializeRoutes() {
	web.Router.HandleFunc("/", web.homeHandler).Methods("GET")
	web.Router.HandleFunc("/login", web.loginHandler).Methods("POST")
	web.Router.HandleFunc("/logout", web.logoutHandler).Methods("GET")
}
