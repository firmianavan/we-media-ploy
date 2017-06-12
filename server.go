package main

import (
	"github.com/firmianavan/we-media-ploy/service"
	"github.com/gorilla/context"
	"log"
	"net/http"
)

func main() {
	mux := service.Mux
	err := http.ListenAndServe(":8080", context.ClearHandler(mux))
	if err != nil {
		log.Fatal("lisent and serve", err)
	}
}
