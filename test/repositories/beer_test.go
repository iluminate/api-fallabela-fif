package repositories

import (
	"api-fallabela-fif/application/entities"
	"api-fallabela-fif/application/repositories"
	"api-fallabela-fif/helpers/database"
	"api-fallabela-fif/test/mocks"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestNewBeerRepository(t *testing.T) {

	mockMongodbHelper := new(mocks.MockMongodbHelper)

	type args struct {
		database database.IMongodbHelper
	}

	tests := []struct {
		name string
		args args
		want repositories.IBeerRepository
	}{
		{
			name: "success",
			args: args{database: mockMongodbHelper},
			want: repositories.NewBeerRepository(mockMongodbHelper)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repositories.NewBeerRepository(tt.args.database); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBeerRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_beerRepository_Create(t *testing.T) {

	sampleBeerSuccess := entities.Beer{
		Id:       1,
		Name:     "Pilsen",
		Brewery:  "Backus",
		Country:  "Peru",
		Price:    5,
		Currency: "PEN",
	}

	mockMongoCollection := new(mocks.MockMongoCollection)
	mockMongoCollection.On("InsertOne", context.Background(), &sampleBeerSuccess).Return(
		&mongo.InsertOneResult{InsertedID: 1}, nil)
	mockMongoCollection.On("InsertOne", context.Background(), &entities.Beer{}).Return(
		&mongo.InsertOneResult{InsertedID: 1}, errors.New("duplicate key error"))

	mockMongodbHelper := new(mocks.MockMongodbHelper)
	mockMongodbHelper.On("Open").Return(nil)
	mockMongodbHelper.On("Collection", "beers").Return(mockMongoCollection)

	type fields struct {
		database database.IMongodbHelper
	}
	type args struct {
		beer *entities.Beer
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "success", fields: fields{database: mockMongodbHelper}, args: args{beer: &sampleBeerSuccess}, wantErr: false},
		{name: "failure", fields: fields{database: mockMongodbHelper}, args: args{beer: &entities.Beer{}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repositories.BeerRepository{
				Database: tt.fields.database,
			}
			if err := repo.Create(tt.args.beer); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_beerRepository_FindAll(t *testing.T) {

	var sampleBeers []entities.Beer

	mockMongoCollection := new(mocks.MockMongoCollection)
	mockMongoCollection.On("Find", context.Background(), bson.D{}).Return(
		&mongo.Cursor{}, errors.New("timeout error"))

	mockMongodbHelper := new(mocks.MockMongodbHelper)
	mockMongodbHelper.On("Open").Return(nil)
	mockMongodbHelper.On("Collection", "beers").Return(mockMongoCollection)

	type fields struct {
		database database.IMongodbHelper
	}

	tests := []struct {
		name    string
		fields  fields
		want    *[]entities.Beer
		wantErr bool
	}{
		{name: "failure", fields: fields{database: mockMongodbHelper}, want: &sampleBeers, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repositories.BeerRepository{
				Database: tt.fields.database,
			}
			got, err := repo.FindAll()
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

func Test_beerRepository_FindById(t *testing.T) {

	mockMongoCollection := new(mocks.MockMongoCollection)
	mockMongoCollection.On("FindOne", context.Background(), bson.D{{"_id", int64(1)}}).Return(
		&mongo.SingleResult{})

	mockMongodbHelper := new(mocks.MockMongodbHelper)
	mockMongodbHelper.On("Open").Return(nil)
	mockMongodbHelper.On("Collection", "beers").Return(mockMongoCollection)

	type fields struct {
		database database.IMongodbHelper
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Beer
		wantErr bool
	}{
		{name: "success", fields: fields{database: mockMongodbHelper}, args: args{id: 1}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repositories.BeerRepository{
				Database: tt.fields.database,
			}
			got, err := repo.FindById(tt.args.id)
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
