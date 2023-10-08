package web

import (
	"log"
	"net/http"
	"text/template"
)

func (web *WebApp) homeHandler(w http.ResponseWriter, r *http.Request) {
	page, ok := pages[r.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	tpl, err := template.ParseFS(res, page)
	if err != nil {
		log.Printf("page %s not found in pages cache...", r.RequestURI)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	data := map[string]interface{}{
		"userAgent": r.UserAgent(),
	}
	if err := tpl.Execute(w, data); err != nil {
		return
	}
}
