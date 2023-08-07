package httpHandler

import (
	"encoding/json"
	"main/src/dbConnect"
	"net/http"
)

func respondWithError(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}

func respondWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		respondWithError(w, "Failed to write JSON data to response", http.StatusInternalServerError)
	}
}

func GetAllTodos(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	todos := dbConnect.GetAllTodos()
	respondWithJSON(w, todos, http.StatusOK)
}

func AddNewTodo(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]string
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		respondWithError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	title, exists := data["title"]
	if !exists {
		respondWithError(w, "Missing 'title' field in JSON", http.StatusBadRequest)
		return
	}

	dbConnect.AddTodo(title)
	respondWithJSON(w, "Todo added successfully", http.StatusCreated)
}

func CheckTodo(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		respondWithError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	idValue, exists := data["id"].(float64)
	if !exists {
		respondWithError(w, "Missing 'id' field in JSON", http.StatusBadRequest)
		return
	}

	idInt := int(idValue)
	dbConnect.CheckTodo(idInt)
	respondWithJSON(w, "Todo updated successfully", http.StatusOK)
}

func DeleteTodo(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		respondWithError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	idValue, exists := data["id"].(float64)
	if !exists {
		respondWithError(w, "Missing 'id' field in JSON", http.StatusBadRequest)
		return
	}

	idInt := int(idValue)
	dbConnect.DeleteTodo(idInt)
	respondWithJSON(w, "Todo removed successfully", http.StatusOK)
}
