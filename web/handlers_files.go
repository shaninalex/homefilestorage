package web

import (
	"net/http"
)

func (web *WebApp) filesUploadHandler(w http.ResponseWriter, r *http.Request) {
	web.helperFileSave(w, r)
}

func (web *WebApp) fileDownloadHandler(w http.ResponseWriter, r *http.Request) {
	web.helperFileGet(w, r)
}

func (web *WebApp) fileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	web.helperFileDelete(w, r)
}

func (web *WebApp) filePreviewHandler(w http.ResponseWriter, r *http.Request) {
	web.helperFilePreview(w, r)
}
