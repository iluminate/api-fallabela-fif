package services

import "api-fallabela-fif/application/models"

type IBeerService interface {
	FindById(id int64) (*models.Beer, error)
	FindAll() (*[]models.Beer, error)
	Create(beer *models.Beer) error
}

type IExchangeService interface {
	Live() (*models.Currency, error)
}
