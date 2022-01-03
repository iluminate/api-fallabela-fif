package services

import (
	"api-fallabela-fif/application/entities"
	"api-fallabela-fif/application/models"
	"api-fallabela-fif/application/repositories"
	"api-fallabela-fif/application/services"
	"api-fallabela-fif/test/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func Test_beerService_Create(t *testing.T) {
	mockBeerRepository := new(mocks.MockBeerRepository)
	mockBeerRepository.On("Create", mock.AnythingOfType("*entities.Beer")).Return(nil)
	type fields struct {
		repo repositories.IBeerRepository
	}
	type args struct {
		beer *models.Beer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "success", fields: fields{
			repo: mockBeerRepository,
		}, args: args{beer: &models.Beer{
			Id:       0,
			Name:     "",
			Brewery:  "",
			Country:  "",
			Price:    0,
			Currency: "",
		}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := services.BeerService{
				Repo: tt.fields.repo,
			}
			if err := service.Create(tt.args.beer); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_beerService_FindAll(t *testing.T) {
	mockBeerRepository := new(mocks.MockBeerRepository)
	mockBeerRepository.On("FindAll").Return(&[]entities.Beer{
		{Id: int64(1), Name: "", Brewery: "", Country: "", Price: float64(0), Currency: ""},
	}, nil)
	type fields struct {
		repo repositories.IBeerRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    *[]models.Beer
		wantErr bool
	}{
		{name: "success", fields: fields{repo: mockBeerRepository},
			want: &[]models.Beer{
				{Id: int64(1), Name: "", Brewery: "", Country: "", Price: float64(0), Currency: ""},
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := services.BeerService{
				Repo: tt.fields.repo,
			}
			got, err := service.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_beerService_FindById(t *testing.T) {
	mockBeerRepository := new(mocks.MockBeerRepository)
	mockBeerRepository.On("FindById", int64(1)).Return(&entities.Beer{}, nil)

	type fields struct {
		repo repositories.IBeerRepository
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Beer
		wantErr bool
	}{
		{name: "success", fields: fields{repo: mockBeerRepository}, args: args{id: int64(1)}, want: &models.Beer{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := services.BeerService{
				Repo: tt.fields.repo,
			}
			got, err := service.FindById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
