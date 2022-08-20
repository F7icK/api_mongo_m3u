package service

import (
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/service/common"
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/service/workers"
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/types"
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/types/config"
	"github.com/F7icK/api_mongo_m3u/internal/clients/repository"
	"github.com/F7icK/api_mongo_m3u/pkg/logger"
	"github.com/pkg/errors"
)

type Common interface {
	GetTracksService() ([]types.Track, error)
	AddTrackService(track *types.PostTrack) error
	DropTrackService(track *types.DeleteTrack) error
	UpdateAllTracks()
}

type Service struct {
	Common
}

func NewService(
	cfg *config.Config,
	db *repository.Repository,
	worker workers.Worker,
) *Service {

	allTracks, err := db.GetTracks()
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with GetTracks in NewService"))
	}

	cc := &Service{
		Common: common.NewCommonService(cfg, db.Common, allTracks),
	}

	if err = worker.AddWorkEveryMinutes(1, cc.UpdateAllTracks); err != nil {
		return nil
	}

	return cc
}
