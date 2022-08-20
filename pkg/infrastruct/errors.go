package infrastruct

import "net/http"

type CustomError struct {
	msg  string
	Code int
}

func NewError(msg string, code int) *CustomError {
	return &CustomError{
		msg:  msg,
		Code: code,
	}
}

func (c *CustomError) Error() string {
	return c.msg
}

var (
	ErrorInternalServerError = NewError("internal server error", http.StatusInternalServerError)
	ErrorBadRequest          = NewError("bad query input", http.StatusBadRequest)
	ErrorBadURI              = NewError("bad query input URI", http.StatusBadRequest)
	ErrorBadName             = NewError("bad query input, missing name", http.StatusBadRequest)
	ErrorTrackExists         = NewError("this track exists", http.StatusConflict)
)
