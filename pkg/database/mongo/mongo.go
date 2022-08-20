package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/types/config"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Database
}

func NewMongoDB(mongoDB *config.MongoDB) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s", mongoDB.Username, mongoDB.Password, mongoDB.URI)))
	if err != nil {
		return nil, errors.Wrap(err, "err with Open DB")
	}

	db := client.Database(mongoDB.DataBase)

	return &MongoDB{Client: db}, nil
}
