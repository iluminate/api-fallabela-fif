package services

import (
	"api-fallabela-fif/application/models"
	"encoding/json"
	"io"
	"net/http"
)

type exchangeService struct {
	Url       string
	Endpoints map[string]string
	Token     string
}

func NewExchangeService(url string, endpoints map[string]string, token string) *exchangeService {
	return &exchangeService{
		Url:       url,
		Endpoints: endpoints,
		Token:     token,
	}
}

func (service exchangeService) Live() (*models.Currency, error) {
	resp, err := http.Get(service.Url + "/" + service.Endpoints["live"])
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var currency models.Currency
	if err = json.Unmarshal(body, &currency); err != nil {
		return nil, err
	}
	return &currency, nil
}
