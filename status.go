package networkutil

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
}

func ErrorStatus(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	response := errorResponse{
		StatusCode: statusCode,
		Status:     err.Error(),
	}
	responseText, err := json.Marshal(response)
	if err != nil {
		ErrorResponse(w, 1, err)
	}

	switch statusCode {
	case 1:
		ErrorStatus(w)
		break
	case 2:
		w.WriteHeader(http.StatusUnauthorized)
		break
	default:
		ErrorStatus(w)
		break
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseText)
}
