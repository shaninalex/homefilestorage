package restapi

import (
	"fmt"
	"net/http"
)

func RouteCreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "Index Route")
}

func RouteIndex(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "")
}
