package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"todo-app/internal/usecases"
)

type FileHandler struct {
	fileService *usecases.FileService
}

func NewFileHandler(fileService *usecases.FileService) *FileHandler {
	return &FileHandler{fileService: fileService}
}

func (h *FileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	if r.ContentLength > 10*1024*1024 {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileID, err := h.fileService.UploadFile(r.Context(), fileBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"fileId": fileID})
}
