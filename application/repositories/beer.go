package repositories

import (
	"api-fallabela-fif/application/entities"
	"api-fallabela-fif/helpers/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

const collectionName = "beers"

type BeerRepository struct {
	Database database.IMongodbHelper
}

func NewBeerRepository(database database.IMongodbHelper) *BeerRepository {
	return &BeerRepository{Database: database}
}

func (repo BeerRepository) FindById(id int64) (*entities.Beer, error) {
	err := repo.Database.Open()
	if err != nil {
		return nil, err
	}
	collection := repo.Database.Collection(collectionName)
	ctx := context.Background()
	var beer entities.Beer
	filters := bson.D{
		{"_id", id},
	}
	err = collection.FindOne(ctx, filters).Decode(&beer)
	if err != nil {
		return nil, err
	}
	return &beer, nil
}

func (repo BeerRepository) FindAll() (*[]entities.Beer, error) {
	var beers []entities.Beer
	err := repo.Database.Open()
	if err != nil {
		return nil, err
	}
	collection := repo.Database.Collection(collectionName)
	ctx := context.Background()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return &beers, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var beer entities.Beer
		err = cur.Decode(&beer)
		if err != nil {
			return nil, err
		}
		beers = append(beers, beer)
	}
	return &beers, nil
}

func (repo BeerRepository) Create(beer *entities.Beer) error {
	err := repo.Database.Open()
	if err != nil {
		return err
	}
	collection := repo.Database.Collection(collectionName)
	ctx := context.Background()
	_, err = collection.InsertOne(ctx, beer)
	if err != nil {
		return err
	}
	return nil
}
