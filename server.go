package main

import (
	"github.com/firmianavan/we-media-ploy/service"
	"log"

	"net/http"
)

func main() {
	mux := service.Mux
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("lisent and serve", err)
	}
}
