package repository

import (
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/types"
	"github.com/F7icK/api_mongo_m3u/internal/clients/repository/common"
	"go.mongodb.org/mongo-driver/mongo"
)

type Common interface {
	GetTracks() ([]types.Track, error)
	InsertTrack(track *types.PostTrack) (*types.Track, error)
	DeleteTrack(track *types.DeleteTrack) error
}

type Repository struct {
	Common
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Common: common.NewCommonMongo(db),
	}
}
