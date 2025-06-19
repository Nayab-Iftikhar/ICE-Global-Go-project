package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"todo-app/internal/usecases"
)

type TodoHandler struct {
	todoService *usecases.TodoService
}

func NewTodoHandler(todoService *usecases.TodoService) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

func (h *TodoHandler) CreateTodoItem(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Description string `json:"description"`
		DueDate     string `json:"dueDate"`
		FileID      string `json:"fileId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dueDate := parseDueDate(req.DueDate)

	todoItem, err := h.todoService.CreateTodoItem(r.Context(), req.Description, dueDate, req.FileID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todoItem)
}

func parseDueDate(dateStr string) time.Time {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		
		return time.Now()
	}
	return t
}
