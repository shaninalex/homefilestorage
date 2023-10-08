package web

import (
	"context"
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/shaninalex/homefilestorage/pkg/config"
	"github.com/shaninalex/homefilestorage/pkg/database"
)

var (
	//go:embed resources
	res   embed.FS
	pages = map[string]string{
		"/":      "resources/index.gohtml",
		"/login": "resources/login.gohtml",
		"__head": "resources/__head.tmpl",
	}
)

type WebApp struct {
	Context  context.Context
	Config   *config.Config
	Database database.Repository
	Router   *mux.Router
}

func InitializeWebApp(conf *config.Config, db database.Repository) (*WebApp, error) {
	webapp := &WebApp{
		Context:  context.TODO(),
		Config:   conf,
		Database: db,
		Router:   mux.NewRouter(),
	}
	http.FileServer(http.FS(res))
	webapp.initializeRoutes()
	return webapp, nil
}

func (web *WebApp) Run(port int64) {
	log.Printf("App started on port %d\n", port)
	// TODO: csrf token authKey config option
	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"))
	web.Router.Use(csrfMiddleware)
	http.ListenAndServe(fmt.Sprintf(":%d", port), web.Router)
}

func (web *WebApp) initializeRoutes() {
	web.Router.HandleFunc("/", web.homeHandler)
	web.Router.HandleFunc("/login", web.loginHandler)
}
