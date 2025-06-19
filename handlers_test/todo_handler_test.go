package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-app/internal/entities"
	"todo-app/internal/handlers"
	"todo-app/internal/usecases"
	"todo-app/mocks"

	"github.com/stretchr/testify/mock"
)

func TestCreateTodoItem_Success(t *testing.T) {
	mockTodoRepo := mocks.NewTodoRepository(t)
	mockFileStorage := mocks.NewFileStorage(t)
	mockRedisStream := mocks.NewRedisStream(t)

	mockTodoRepo.On("Save", mock.Anything, mock.Anything).Return(nil).Once()
	mockRedisStream.On("Publish", mock.Anything, mock.Anything).Return(nil).Once()

	todoService := usecases.NewTodoService(mockTodoRepo, mockFileStorage, mockRedisStream)
	handler := handlers.NewTodoHandler(todoService)

	reqBody := `{
		"description": "Test Todo",
		"dueDate": "2025-06-19T12:00:00Z",
		"fileId": "file-id-123"
	}`
	req := httptest.NewRequest("POST", "/todo", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.CreateTodoItem(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rr.Code)
	}

	var todoItem entities.TodoItem
	if err := json.NewDecoder(rr.Body).Decode(&todoItem); err != nil {
		t.Fatal(err)
	}

	if todoItem.Description != "Test Todo" {
		t.Fatalf("expected description 'Test Todo', got '%s'", todoItem.Description)
	}

	mockTodoRepo.AssertExpectations(t)
}
