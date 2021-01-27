package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON write status code and data into the response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	if error := json.NewEncoder(w).Encode(data); error != nil {
		log.Fatal(error)
	}
}

// Error write status code and error into the response
func Error(w http.ResponseWriter, statusCode int, error error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: error.Error(),
	})
}
