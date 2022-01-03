package mocks

import (
	"api-fallabela-fif/application/entities"
	"api-fallabela-fif/application/models"
	"github.com/stretchr/testify/mock"
	"time"
)

type MockContext struct {
	mock.Mock
}

func (MockContext) Deadline() (deadline time.Time, ok bool) {
	//TODO implement me
	panic("implement me")
}

func (MockContext) Done() <-chan struct{} {
	//TODO implement me
	panic("implement me")
}

func (MockContext) Err() error {
	//TODO implement me
	panic("implement me")
}

func (m *MockContext) Value(key interface{}) interface{} {
	args := m.Called(key)
	return args.Get(0)
}

type MockBeerService struct {
	mock.Mock
}

func (m *MockBeerService) FindById(id int64) (*models.Beer, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Beer), args.Error(1)
	} else {
		return nil, args.Error(1)
	}

}

func (m *MockBeerService) FindAll() (*[]models.Beer, error) {
	args := m.Called()
	return args.Get(0).(*[]models.Beer), args.Error(1)
}

func (m *MockBeerService) Create(beer *models.Beer) error {
	args := m.Called(beer)
	return args.Error(0)
}

type MockExchangeService struct {
	mock.Mock
}

func (m *MockExchangeService) Live() (*models.Currency, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).(*models.Currency), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

type MockBeerRepository struct {
	mock.Mock
}

func (m *MockBeerRepository) FindById(id int64) (*entities.Beer, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Beer), args.Error(1)
}

func (m *MockBeerRepository) FindAll() (*[]entities.Beer, error) {
	args := m.Called()
	return args.Get(0).(*[]entities.Beer), args.Error(1)
}

func (m *MockBeerRepository) Create(beer *entities.Beer) error {
	args := m.Called(beer)
	return args.Error(0)
}
