package handlers

import (
	"net/http"

	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/server/handlers/common"
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/service"
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/types/config"
)

type Common interface {
	Ping(w http.ResponseWriter, _ *http.Request)
	GetPlayList(w http.ResponseWriter, _ *http.Request)
	PostTrack(w http.ResponseWriter, r *http.Request)
	DeleteTrack(w http.ResponseWriter, r *http.Request)
}

type Handlers struct {
	Common
}

func NewHandlers(s *service.Service, cfg *config.Config) *Handlers {
	return &Handlers{
		Common: common.NewCommonHandlers(s),
	}
}
