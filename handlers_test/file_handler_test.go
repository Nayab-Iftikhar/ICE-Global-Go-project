package handlers_test

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-app/internal/handlers"
	"todo-app/internal/usecases"
	"todo-app/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/mock"
)

func TestUploadFile_Success(t *testing.T) {
	
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockFileStorage := mocks.NewFileStorage(t)

	mockFileStorage.On("UploadFile", mock.Anything, mock.Anything).Return("file-id-123", nil)

	fileService := usecases.NewFileService(mockFileStorage)
	handler := handlers.NewFileHandler(fileService)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("file", "test.txt")
	if err != nil {
		t.Fatal(err)
	}
	fileContent := []byte("dummy content")
	file.Write(fileContent)
	writer.Close()

	t.Logf("Uploading file: name=%s, content=%s", "test.txt", string(fileContent))

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()

	handler.UploadFile(rr, req)

	t.Logf("Response status: %d", rr.Code)
	t.Logf("Response body: %s", rr.Body.String())

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}
	if response["fileId"] != "file-id-123" {
		t.Fatalf("expected fileId 'file-id-123', got '%s'", response["fileId"])
	}
}
