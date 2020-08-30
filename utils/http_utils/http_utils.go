package http_utils

import (
	"encoding/json"
	"github.com/JingdaMai/bookstore_utils-go/rest_errors"
	"log"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Printf("Error writing json: %v\n", err)
	}
}

func RespondError(w http.ResponseWriter, err rest_errors.RestErr) {
	RespondJson(w, err.Status(), err)
}
