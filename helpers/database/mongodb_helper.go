package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongodbHelper struct {
	client *mongo.Client
	conf   map[string]string
}

func NewMongodbHelper(conf map[string]string) *MongodbHelper {
	return &MongodbHelper{
		conf: conf,
	}
}
func (helper *MongodbHelper) Collection(name string) *mongo.Collection {
	return helper.client.Database(helper.conf["database"]).Collection(name)
}

func (helper *MongodbHelper) Open() error {
	var err error
	if helper.client != nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	helper.client, err = mongo.Connect(ctx, options.Client().ApplyURI(helper.conf["uri"]))
	if err != nil {
		return err
	}
	err = helper.client.Ping(context.TODO(), nil)
	if err != nil {
		defer helper.close()
		return err
	}
	log.Println("connected to mongo!")
	return nil
}

func (helper *MongodbHelper) close() error {
	err := helper.client.Disconnect(context.TODO())
	if err != nil {
		return err
	}
	fmt.Println("connected to mongo closed.")
	return err
}
