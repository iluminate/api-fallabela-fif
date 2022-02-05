package services

import (
	"api-fallabela-fif/application/models"
	"api-fallabela-fif/application/services"
	"api-fallabela-fif/test/mocks"
	"net/http"
	"reflect"
	"testing"
)

func Test_exchangeService_Live(t *testing.T) {

	currencies := new(models.Currency)
	serverMock := mocks.MockServerWithOut(currencies)

	defer serverMock.Close()

	type fields struct {
		Url       string
		Endpoints map[string]string
		Token     string
		Client    *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *models.Currency
		wantErr bool
	}{
		{name: "success", fields: fields{
			Url:       serverMock.URL,
			Endpoints: map[string]string{"live": "live"},
			Token:     "41828eb46c449c80a074234172fa9498",
			Client:    serverMock.Client(),
		}, want: currencies, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := services.ExchangeService{
				Url:       tt.fields.Url,
				Endpoints: tt.fields.Endpoints,
				Token:     tt.fields.Token,
				Client:    tt.fields.Client,
			}
			got, err := service.Live()
			if (err != nil) != tt.wantErr {
				t.Errorf("Live() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Live() got = %v, want %v", got, tt.want)
			}
		})
	}
}
