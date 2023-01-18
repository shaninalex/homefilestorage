package restapi

import (
	"net/http"

	"gorm.io/gorm"
)

func Server(db *gorm.DB) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", RouteIndex)
	mux.HandleFunc("/api/v1/account/create/", RouteCreateUser)

	return
}
