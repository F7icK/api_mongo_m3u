package encode

import (
	"encoding/json"
	"net/http"

	"github.com/F7icK/api_mongo_m3u/pkg/infrastruct"
	"github.com/F7icK/api_mongo_m3u/pkg/logger"
)

type Result struct {
	Err string `json:"error"`
}

func ApiErrorEncode(w http.ResponseWriter, err error) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if customError, ok := err.(*infrastruct.CustomError); ok {
		w.WriteHeader(customError.Code)
	}

	r := Result{Err: err.Error()}

	if err = json.NewEncoder(w).Encode(r); err != nil {
		logger.LogError(err)
	}
}

func ApiResponseEncoder(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		logger.LogError(err)
	}
}

func ApiResponseEncoderStatusCode(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
}
