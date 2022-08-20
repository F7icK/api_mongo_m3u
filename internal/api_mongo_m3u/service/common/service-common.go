package common

import (
	"net/url"

	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/types"
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/types/config"
	"github.com/F7icK/api_mongo_m3u/internal/clients/repository"
	"github.com/F7icK/api_mongo_m3u/pkg/infrastruct"
	"github.com/F7icK/api_mongo_m3u/pkg/logger"
	"github.com/pkg/errors"
)

type CommonService struct {
	cfg       *config.Config
	db        repository.Common
	allTracks []types.Track
}

func NewCommonService(
	cfg *config.Config,
	db repository.Common,
	allTracks []types.Track,
) *CommonService {
	return &CommonService{
		cfg:       cfg,
		db:        db,
		allTracks: allTracks,
	}
}

func (s *CommonService) GetTracksService() ([]types.Track, error) {
	tracks, err := s.db.GetTracks()
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with GetTracks in GetTracksService"))
		return nil, infrastruct.ErrorInternalServerError
	}

	return tracks, nil
}

func (s *CommonService) AddTrackService(track *types.PostTrack) error {
	if err := validateTrack(track); err != nil {
		return err
	}

	for _, dbTrack := range s.allTracks {
		if dbTrack.URI == track.URI && dbTrack.Name == track.Name {
			return infrastruct.ErrorTrackExists
		}
	}

	newTrack, err := s.db.InsertTrack(track)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with InsertTrack in AddTrackService"))
		return infrastruct.ErrorInternalServerError
	}

	s.allTracks = append(s.allTracks, *newTrack)

	return nil
}

func (s *CommonService) DropTrackService(track *types.DeleteTrack) error {
	if err := s.db.DeleteTrack(track); err != nil {
		logger.LogError(errors.Wrap(err, "err with DeleteTrack in DropTrackService"))
		return infrastruct.ErrorInternalServerError
	}

	for idx, dbTrack := range s.allTracks {
		if dbTrack.Id == track.Id {
			s.allTracks[idx] = s.allTracks[len(s.allTracks)-1]
			s.allTracks = s.allTracks[:len(s.allTracks)-1]
			break
		}
	}

	return nil
}

func validateTrack(track *types.PostTrack) error {
	if _, err := url.ParseRequestURI(track.URI); err != nil {
		return infrastruct.ErrorBadURI
	}

	if track.Name == "" {
		return infrastruct.ErrorBadName
	}

	return nil
}

func (s *CommonService) UpdateAllTracks() {
	tracks, err := s.db.GetTracks()
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with GetTracks in GetTracksService"))
		return
	}

	s.allTracks = tracks
}
