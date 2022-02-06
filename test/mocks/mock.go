package mocks

import (
	"api-fallabela-fif/application/entities"
	"api-fallabela-fif/application/models"
	"api-fallabela-fif/helpers/database"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"net/http/httptest"
	"time"
)

type MockContext struct {
	mock.Mock
}

func (m *MockContext) Deadline() (deadline time.Time, ok bool) {
	args := m.Called()
	return args.Get(0).(time.Time), args.Bool(1)
}

func (m *MockContext) Done() <-chan struct{} {
	return nil
}

func (m *MockContext) Err() error {
	args := m.Called()
	return args.Error(0)
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

func MockServerWithOut(response interface{}) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		out, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		w.Write(out)
	}))
}

type MockMongodbHelper struct {
	mock.Mock
}

func (m *MockMongodbHelper) Collection(name string) database.IMongoCollection {
	args := m.Called(name)
	return args.Get(0).(database.IMongoCollection)
}

func (m *MockMongodbHelper) Open() error {
	args := m.Called()
	return args.Error(0)
}

type MockMongoCollection struct {
	mock.Mock
}

func (m *MockMongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockMongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *MockMongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}
