package repositories

import (
	"api-fallabela-fif/application/entities"
	"api-fallabela-fif/application/repositories"
	"api-fallabela-fif/helpers/database"
	"reflect"
	"testing"
)

func TestNewBeerRepository(t *testing.T) {
	type args struct {
		database *database.MongodbHelper
	}
	tests := []struct {
		name string
		args args
		want *repositories.beerRepository
	}{
		// TODO: Add test cases.
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
	type fields struct {
		database *database.MongodbHelper
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repositories.beerRepository{
				database: tt.fields.database,
			}
			if err := repo.Create(tt.args.beer); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_beerRepository_FindAll(t *testing.T) {
	type fields struct {
		database *database.MongodbHelper
	}
	tests := []struct {
		name    string
		fields  fields
		want    *[]entities.Beer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repositories.beerRepository{
				database: tt.fields.database,
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
	type fields struct {
		database *database.MongodbHelper
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repositories.beerRepository{
				database: tt.fields.database,
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
