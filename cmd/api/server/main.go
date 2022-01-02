package main

import (
	"api-fallabela-fif/application/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.NewHttpRouter()
	server := http.Server{
		Handler: router.Handler(),
		Addr:    ":8080",
	}
	log.Fatalln(server.ListenAndServe())
}
