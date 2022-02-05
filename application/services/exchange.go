package services

import (
	"api-fallabela-fif/application/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ExchangeService struct {
	Url       string
	Endpoints map[string]string
	Token     string
	Client    *http.Client
}

func NewExchangeService(url string, endpoints map[string]string, token string, client *http.Client) *ExchangeService {
	return &ExchangeService{
		Url:       url,
		Endpoints: endpoints,
		Token:     token,
		Client:    client,
	}
}

func (service ExchangeService) Live() (*models.Currency, error) {
	resp, err := service.Client.Get(service.Url + "/" + service.Endpoints["live"])
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	var currency models.Currency
	if err = json.Unmarshal(body, &currency); err != nil {
		return nil, err
	}
	return &currency, nil
}
