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

	router.GET("/", RouteIndex)
	router.POST("/api/v1/account/create/", RouteCreateUser)

	log.Printf("Start server under :%d port...", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
