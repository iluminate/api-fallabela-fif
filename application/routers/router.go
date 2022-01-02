package routers

import (
	"api-fallabela-fif/application/containers"
	"api-fallabela-fif/application/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

type httpRouter struct {
	beerHandler *controllers.BeerHandler
}

func NewHttpRouter() *httpRouter {
	return &httpRouter{beerHandler: containers.BeerHandler()}
}

func (httpRouter httpRouter) Handler() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", baseHandler).Methods("GET")
	router.HandleFunc("/beers", httpRouter.beerHandler.GetBeers).Methods("GET")
	router.HandleFunc("/beers", httpRouter.beerHandler.PostBeers).Methods("POST")
	router.HandleFunc("/beers/{id}", httpRouter.beerHandler.GetBeersPerId).Methods("GET")
	router.HandleFunc("/beers/{id}/boxprice", httpRouter.beerHandler.GetBoxPrixePerId).Methods("GET")
	return router
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("api is up!"))
}
