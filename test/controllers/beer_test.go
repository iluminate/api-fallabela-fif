package controllers

import (
	"api-fallabela-fif/application/controllers"
	"api-fallabela-fif/application/models"
	"api-fallabela-fif/application/services"
	"api-fallabela-fif/test/mocks"
	"errors"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBeerHandler_GetBeers(t *testing.T) {
	mockBeerService := new(mocks.MockBeerService)
	mockBeerService.On("FindAll").Return(&[]models.Beer{}, nil)
	type fields struct {
		beerService     services.IBeerService
		exchangeService services.IExchangeService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "success", fields: fields{
			beerService:     mockBeerService,
			exchangeService: nil,
		}, args: args{
			w: httptest.NewRecorder(),
			r: makeHttpRequest("GET", "/beers", nil, nil),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &controllers.BeerHandler{
				BeerService:     tt.fields.beerService,
				ExchangeService: tt.fields.exchangeService,
			}
			handler.GetBeers(tt.args.w, tt.args.r)
		})
	}
}

func TestBeerHandler_GetBeersPerId(t *testing.T) {
	mockBeerService := new(mocks.MockBeerService)
	mockBeerService.On("FindById", int64(1)).Return(&models.Beer{}, nil)
	mockBeerService.On("FindById", int64(2)).Return(nil, errors.New("no document"))
	type fields struct {
		beerService     services.IBeerService
		exchangeService services.IExchangeService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "success", fields: fields{
			beerService:     mockBeerService,
			exchangeService: nil,
		}, args: args{
			w: httptest.NewRecorder(),
			r: makeHttpRequest("GET", "/beers/1", map[string]string{"id": "1"}, nil),
		}},
		{name: "failure", fields: fields{
			beerService:     mockBeerService,
			exchangeService: nil,
		}, args: args{
			w: httptest.NewRecorder(),
			r: makeHttpRequest("GET", "/beers/2", map[string]string{"id": "2"}, nil),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &controllers.BeerHandler{
				BeerService:     tt.fields.beerService,
				ExchangeService: tt.fields.exchangeService,
			}
			handler.GetBeersPerId(tt.args.w, tt.args.r)
		})
	}
}

func TestBeerHandler_GetBoxPrixePerId(t *testing.T) {
	mockBeerService := new(mocks.MockBeerService)
	mockBeerService.On("FindById", int64(1)).Return(&models.Beer{}, nil)
	mockBeerService.On("FindById", int64(2)).Return(nil, errors.New("no document"))
	mockExchangeService := new(mocks.MockExchangeService)
	mockExchangeService.On("Live").Return(&models.Currency{}, nil)
	mockExchangeErrorService := new(mocks.MockExchangeService)
	mockExchangeErrorService.On("Live").Return(nil, errors.New("service error"))
	type fields struct {
		beerService     services.IBeerService
		exchangeService services.IExchangeService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "success", fields: fields{
			beerService:     mockBeerService,
			exchangeService: mockExchangeService,
		}, args: args{
			w: httptest.NewRecorder(),
			r: makeHttpRequest("GET", "/beers/1/boxprice", map[string]string{"id": "1"}, nil),
		}},
		{name: "failure1", fields: fields{
			beerService:     mockBeerService,
			exchangeService: mockExchangeService,
		}, args: args{
			w: httptest.NewRecorder(),
			r: makeHttpRequest("GET", "/beers/2/boxprice", map[string]string{"id": "2"}, nil),
		}},
		{name: "failure2", fields: fields{
			beerService:     mockBeerService,
			exchangeService: mockExchangeErrorService,
		}, args: args{
			w: httptest.NewRecorder(),
			r: makeHttpRequest("GET", "/beers/1/boxprice", map[string]string{"id": "1"}, nil),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &controllers.BeerHandler{
				BeerService:     tt.fields.beerService,
				ExchangeService: tt.fields.exchangeService,
			}
			handler.GetBoxPrixePerId(tt.args.w, tt.args.r)
		})
	}
}

func TestBeerHandler_PostBeers(t *testing.T) {
	mockBeerService := new(mocks.MockBeerService)
	mockBeerService.On("Create", &models.Beer{
		Id:       2,
		Name:     "Pilsen",
		Brewery:  "unkwoun",
		Country:  "Peru",
		Price:    6.5,
		Currency: "USD",
	}).Return(errors.New("id dup key"))
	mockBeerService.On("Create", &models.Beer{
		Id:       1,
		Name:     "Pilsen",
		Brewery:  "unkwoun",
		Country:  "Peru",
		Price:    6.5,
		Currency: "USD",
	}).Return(nil)
	type fields struct {
		beerService     services.IBeerService
		exchangeService services.IExchangeService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "success", fields: fields{
			beerService:     mockBeerService,
			exchangeService: nil,
		}, args: args{
			w: httptest.NewRecorder(),
			r: makeHttpRequest("POST", "/beers", nil, strings.NewReader("{\n    \"id\": 1,\n    \"name\": \"Pilsen\",\n    \"brewery\": \"unkwoun\",\n    \"country\": \"Peru\",\n    \"price\": 6.5,\n    \"currency\": \"USD\"\n}")),
		}},
		{name: "failure1", fields: fields{
			beerService:     mockBeerService,
			exchangeService: nil,
		}, args: args{
			w: httptest.NewRecorder(),
			r: makeHttpRequest("POST", "/beers", nil, strings.NewReader("{\n    \"id\": \"2\",\n    \"name\": \"Pilsen\",\n    \"brewery\": \"unkwoun\",\n    \"country\": \"Peru\",\n    \"price\": 6.5,\n    \"currency\": \"USD\"\n}")),
		}},
		{name: "failure2", fields: fields{
			beerService:     mockBeerService,
			exchangeService: nil,
		}, args: args{
			w: httptest.NewRecorder(),
			r: makeHttpRequest("POST", "/beers", nil, strings.NewReader("{\n    \"id\": 2,\n    \"name\": \"Pilsen\",\n    \"brewery\": \"unkwoun\",\n    \"country\": \"Peru\",\n    \"price\": 6.5,\n    \"currency\": \"USD\"\n}")),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &controllers.BeerHandler{
				BeerService:     tt.fields.beerService,
				ExchangeService: tt.fields.exchangeService,
			}
			handler.PostBeers(tt.args.w, tt.args.r)
		})
	}
}

func makeHttpRequest(method string, uri string, value map[string]string, body io.Reader) *http.Request {
	mockContext := new(mocks.MockContext)
	mockContext.On("Value", mock.AnythingOfType("contextKey")).Return(value)
	r, _ := http.NewRequestWithContext(mockContext, method, uri, body)
	return r
}
