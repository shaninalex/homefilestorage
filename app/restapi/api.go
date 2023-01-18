package restapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
	"gorm.io/gorm"
)

func Server(db *gorm.DB, port int) {

	router := bunrouter.New()

	h := Handlers(db)

	router.GET("/", h.RouteIndex)
	router.POST("/api/v1/account/create/", h.RouteCreateUser)

	log.Printf("Start server under :%d port...", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
