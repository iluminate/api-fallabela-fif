package services

import (
	"api-fallabela-fif/application/entities"
	"api-fallabela-fif/application/models"
	"api-fallabela-fif/application/repositories"
)

type BeerService struct {
	Repo repositories.IBeerRepository
}

func NewBeerService(repo repositories.IBeerRepository) *BeerService {
	return &BeerService{Repo: repo}
}

func (service BeerService) FindById(id int64) (*models.Beer, error) {
	beer, err := service.Repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return &models.Beer{
		Id:       beer.Id,
		Name:     beer.Name,
		Brewery:  beer.Brewery,
		Country:  beer.Country,
		Price:    beer.Price,
		Currency: beer.Currency,
	}, nil
}

func (service BeerService) FindAll() (*[]models.Beer, error) {
	beers, err := service.Repo.FindAll()
	if err != nil {
		return nil, err
	}
	res := []models.Beer{}
	for _, beer := range *beers {
		res = append(res, models.Beer{
			Id:       beer.Id,
			Name:     beer.Name,
			Brewery:  beer.Brewery,
			Country:  beer.Country,
			Price:    beer.Price,
			Currency: beer.Currency,
		})
	}
	return &res, nil
}

func (service BeerService) Create(beer *models.Beer) error {
	return service.Repo.Create(&entities.Beer{
		Id:       beer.Id,
		Name:     beer.Name,
		Brewery:  beer.Brewery,
		Country:  beer.Country,
		Price:    beer.Price,
		Currency: beer.Currency,
	})
}
