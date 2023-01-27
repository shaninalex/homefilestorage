package restapi

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func Server(db *sql.DB, port int) {

	router := bunrouter.New()

	h := Handlers(db)

	router.GET("/", h.RouteIndex)
	router.POST("/api/v1/account/create/", h.RouteCreateUser)

	log.Printf("Start server under :%d port...", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
