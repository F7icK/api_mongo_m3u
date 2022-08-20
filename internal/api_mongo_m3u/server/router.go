package server

import (
	"net/http"

	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/server/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(h *handlers.Handlers) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.Methods(http.MethodGet).Path("/ping").HandlerFunc(h.Ping)

	router.Methods(http.MethodGet).Path("/action").HandlerFunc(h.GetPlayList)
	router.Methods(http.MethodPost).Path("/action").HandlerFunc(h.PostTrack)
	router.Methods(http.MethodDelete).Path("/action").HandlerFunc(h.DeleteTrack)

	return router
}
