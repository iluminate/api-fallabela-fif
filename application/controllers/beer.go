package controllers

import (
	"api-fallabela-fif/application/models"
	"api-fallabela-fif/application/services"
	"api-fallabela-fif/application/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BeerHandler struct {
	beerService     services.IBeerService
	exchangeService services.IExchangeService
}

func NewBeerHandler(service services.IBeerService, exchangeService services.IExchangeService) *BeerHandler {
	return &BeerHandler{beerService: service, exchangeService: exchangeService}
}

func (handler *BeerHandler) GetBeers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	beers, _ := handler.beerService.FindAll()
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(beers)
}

func (handler *BeerHandler) PostBeers(w http.ResponseWriter, r *http.Request) {
	var beer models.Beer
	err := json.NewDecoder(r.Body).Decode(&beer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request"))
		return
	}
	err = handler.beerService.Create(&beer)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("beer id alreay exist"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("beer created"))
}

func (handler *BeerHandler) GetBeersPerId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	_id, _ := strconv.ParseInt(id, 10, 64)
	beer, err := handler.beerService.FindById(_id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("beer not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(beer)
}

func (handler *BeerHandler) GetBoxPrixePerId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	_id, _ := strconv.ParseInt(id, 10, 64)
	beer, err := handler.beerService.FindById(_id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("beer not found"))
		return
	}
	exchange, err := handler.exchangeService.Live()
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("exchange service error"))
		return
	}
	oldCurrency := beer.Currency
	newCurrency := r.URL.Query().Get("currency")
	amountStr := r.URL.Query().Get("quantity")
	var amountInt, priceExchange float64
	amountInt = utils.ParseFloat(amountStr, amountInt)
	totalPrice := beer.Price * amountInt
	if newCurrency != "" {
		priceExchange = (totalPrice / exchange.Quotes["USD"+oldCurrency]) * exchange.Quotes["USD"+newCurrency]
	} else {
		priceExchange = totalPrice
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(map[string]float64{"Price Total": priceExchange})
}
