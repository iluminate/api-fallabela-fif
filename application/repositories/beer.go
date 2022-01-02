package repositories

import (
	"api-fallabela-fif/application/entities"
	"api-fallabela-fif/helpers/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type beerRepository struct {
	database *database.MongodbHelper
}

func NewBeerRepository(database *database.MongodbHelper) *beerRepository {
	return &beerRepository{database: database}
}

func (repo beerRepository) FindById(id int64) (*entities.Beer, error) {
	err := repo.database.Open()
	if err != nil {
		return nil, err
	}
	collection := repo.database.Collection("beers")
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

func (repo beerRepository) FindAll() (*[]entities.Beer, error) {
	err := repo.database.Open()
	if err != nil {
		return nil, err
	}
	collection := repo.database.Collection("beers")
	ctx := context.Background()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var beers []entities.Beer
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

func (repo beerRepository) Create(beer *entities.Beer) error {
	err := repo.database.Open()
	if err != nil {
		return err
	}
	collection := repo.database.Collection("beers")
	ctx := context.Background()
	_, err = collection.InsertOne(ctx, beer)
	if err != nil {
		return err
	}
	return nil
}
