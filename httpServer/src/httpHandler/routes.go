package httpHandler

import (
	"src/dbConnect"
	"encoding/json"
	"net/http"
)

func GetAllTodos(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write(dbConnect.GetAllTodos())
	if err != nil {
		http.Error(w, "Failed to write JSON data to response", http.StatusInternalServerError)
		return
	}
}
