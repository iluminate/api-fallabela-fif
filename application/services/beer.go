package services

import (
	"api-fallabela-fif/application/entities"
	"api-fallabela-fif/application/models"
	"api-fallabela-fif/application/repositories"
)

type beerService struct {
	repo repositories.IBeerRepository
}

func NewBeerService(repo repositories.IBeerRepository) *beerService {
	return &beerService{repo: repo}
}

func (service beerService) FindById(id int64) (*models.Beer, error) {
	beer, err := service.repo.FindById(id)
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

func (service beerService) FindAll() (*[]models.Beer, error) {
	beers, err := service.repo.FindAll()
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

func (service beerService) Create(beer *models.Beer) error {
	return service.repo.Create(&entities.Beer{
		Id:       beer.Id,
		Name:     beer.Name,
		Brewery:  beer.Brewery,
		Country:  beer.Country,
		Price:    beer.Price,
		Currency: beer.Currency,
	})
}
