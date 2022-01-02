package repositories

import (
	"api-fallabela-fif/application/entities"
)

type IBeerRepository interface {
	FindById(id int64) (*entities.Beer, error)
	FindAll() (*[]entities.Beer, error)
	Create(beer *entities.Beer) error
}
