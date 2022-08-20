package common

import (
	"encoding/json"
	"net/http"

	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/server/handlers/encode"
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/service"
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/types"
	"github.com/F7icK/api_mongo_m3u/pkg/infrastruct"
)

type CommonHandlers struct {
	s service.Common
}

func NewCommonHandlers(s service.Common) *CommonHandlers {
	return &CommonHandlers{
		s: s,
	}
}

func (h *CommonHandlers) Ping(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

func (h *CommonHandlers) GetPlayList(w http.ResponseWriter, _ *http.Request) {
	newPlaylist, err := h.s.GetTracksService()
	if err != nil {
		encode.ApiErrorEncode(w, err)
		return
	}

	encode.ApiResponseEncoder(w, newPlaylist)
}

func (h *CommonHandlers) PostTrack(w http.ResponseWriter, r *http.Request) {
	var req types.PostTrack
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		encode.ApiErrorEncode(w, infrastruct.ErrorBadRequest)
		return
	}

	if err := h.s.AddTrackService(&req); err != nil {
		encode.ApiErrorEncode(w, err)
		return
	}

	encode.ApiResponseEncoderStatusCode(w, http.StatusOK)
}

func (h *CommonHandlers) DeleteTrack(w http.ResponseWriter, r *http.Request) {
	var req types.DeleteTrack
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		encode.ApiErrorEncode(w, infrastruct.ErrorBadRequest)
		return
	}

	if err := h.s.DropTrackService(&req); err != nil {
		encode.ApiErrorEncode(w, err)
		return
	}

	encode.ApiResponseEncoderStatusCode(w, http.StatusOK)
}
