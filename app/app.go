package app

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "index")
}
