package web

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/shaninalex/homefilestorage/pkg/config"
	"github.com/shaninalex/homefilestorage/pkg/database"
)

var (
	//go:embed resources
	res   embed.FS
	pages = map[string]string{
		"/": "resources/index.gohtml",
	}
)

type WebApp struct {
	Config   *config.Config
	Database database.Repository
}

func InitializeWebApp(conf *config.Config, db database.Repository) (*WebApp, error) {
	webapp := &WebApp{
		Config:   conf,
		Database: db,
	}


	http.FileServer(http.FS(res))
	webapp.initializeRoutes()

	return webapp, nil
}

func (web *WebApp) Run(port int64) {
	log.Printf("App started on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func (web *WebApp) initializeRoutes() {

	http.HandleFunc("/", web.homeHandler)
}
