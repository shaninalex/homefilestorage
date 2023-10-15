package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/shaninalex/homefilestorage/pkg/config"
	"github.com/shaninalex/homefilestorage/pkg/database"
	"github.com/shaninalex/homefilestorage/pkg/filemanager"
	"github.com/shaninalex/homefilestorage/web/templates"
)

type WebApp struct {
	Context     context.Context
	Config      *config.Config
	Database    database.Repository
	Router      *mux.Router
	State       *templates.State
	Store       *sessions.CookieStore
	FileManager *filemanager.FileManager
}

func InitializeWebApp(conf *config.Config, db database.Repository) (*WebApp, error) {
	fm, err := filemanager.InitFileManager(conf.Storage.Path)
	if err != nil {
		return nil, err
	}
	webapp := &WebApp{
		Context:     context.TODO(),
		Config:      conf,
		Database:    db,
		Router:      mux.NewRouter(),
		State:       &templates.State{LoggedIn: false},
		Store:       sessions.NewCookieStore([]byte(conf.Web.SecretKey)),
		FileManager: fm,
	}
	err = db.Migrate()
	if err != nil {
		return nil, err
	}
	go webapp.createAdmin()
	go webapp.createStorageFolder()
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
	web.Router.HandleFunc("/files/upload", web.filesUploadHandler).Methods("POST")
	web.Router.HandleFunc("/files/{id:[0-9]+}/download", web.fileDownloadHandler).Methods("GET")
	web.Router.HandleFunc("/files/{id:[0-9]+}/preview", web.filePreviewHandler).Methods("GET")
	web.Router.HandleFunc("/files/{id:[0-9]+}/delete", web.fileDeleteHandler).Methods("POST")
}

func (web *WebApp) createAdmin() error {
	_, err := web.Database.GetAccountByEmail(web.Config.Admin.Email)
	if err == nil {
		// Account existss
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

func (web *WebApp) createStorageFolder() {
	err := os.MkdirAll(web.Config.Storage.Path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
