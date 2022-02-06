package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type IMongodbHelper interface {
	Collection(name string) IMongoCollection
	Open() error
}

type IMongoCollection interface {
	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{},
		opts ...*options.FindOptions) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter interface{},
		opts ...*options.FindOneOptions) *mongo.SingleResult
}

type MongodbHelper struct {
	Client *mongo.Client
	Conf   map[string]string
}

func NewMongodbHelper(conf map[string]string) *MongodbHelper {
	return &MongodbHelper{
		Conf: conf,
	}
}

func (helper *MongodbHelper) Collection(name string) IMongoCollection {
	return helper.Client.Database(helper.Conf["database"]).Collection(name)
}

func (helper *MongodbHelper) Open() error {
	var err error
	if helper.Client != nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	helper.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(helper.Conf["uri"]))
	if err != nil {
		return err
	}
	err = helper.Client.Ping(context.TODO(), nil)
	if err != nil {
		defer helper.close()
		return err
	}
	log.Println("connected to mongo!")
	return nil
}

func (helper *MongodbHelper) close() error {
	err := helper.Client.Disconnect(context.TODO())
	if err != nil {
		return err
	}
	fmt.Println("connected to mongo closed.")
	return err
}
