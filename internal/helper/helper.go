package helper

import (
	"encoding/json"
	"net/http"
)

// WriteCustomResp write custom response http.
func WriteCustomResp(w http.ResponseWriter, headerStatus int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(headerStatus)
	json.NewEncoder(w).Encode(response)
	return
}
