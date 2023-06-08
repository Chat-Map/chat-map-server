package api

import (
	"encoding/json"
	"net/http"

	"github.com/lordvidex/errs"
)

type Response struct {
	// Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Err     error       `json:"error,omitempty"`
}

func (r *Response) Write(w http.ResponseWriter) error {
	if r.Err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(errorStatusCode(r.Err))
	}
	return json.NewEncoder(w).Encode(r)
}

func newSuccessResponse(message string, data interface{}) *Response {
	return &Response{
		// Success: true,
		Message: message,
		Data:    data,
	}
}

func newFailureResponse(message string, err error) *Response {
	return &Response{
		// Success: false,
		Message: message,
		Err:     err,
	}
}

func errorStatusCode(err error) int {
	if err == nil {
		return 200
	}
	if errr, ok := (err).(*errs.Error); ok {
		return errr.Code.HTTP()
	}
	return 500
}
