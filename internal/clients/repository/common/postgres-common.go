package common

import (
	"context"
	"time"

	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommonMongo struct {
	db *mongo.Database
}

func NewCommonMongo(db *mongo.Database) *CommonMongo {
	return &CommonMongo{db: db}
}

func (m *CommonMongo) GetTracks() ([]types.Track, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	tracks := make([]types.Track, 0)
	cur, err := m.db.Collection("tracks").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, &tracks); err != nil {
		return nil, err
	}

	return tracks, nil
}

func (m *CommonMongo) InsertTrack(track *types.PostTrack) (*types.Track, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := m.db.Collection("tracks").InsertOne(ctx, track)
	if err != nil {
		return nil, err
	}

	newTrack := &types.Track{
		Id:   res.InsertedID.(primitive.ObjectID),
		Name: track.Name,
		URI:  track.URI,
	}

	return newTrack, nil
}

func (m *CommonMongo) DeleteTrack(track *types.DeleteTrack) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := m.db.Collection("tracks").DeleteOne(ctx, track); err != nil {
		return err
	}

	return nil
}
