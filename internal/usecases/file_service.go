package usecases

import (
	"context"
	"todo-app/internal/interfaces"
)

type FileService struct {
	fileStorage interfaces.FileStorage
}

func NewFileService(fileStorage interfaces.FileStorage) *FileService {
	return &FileService{
		fileStorage: fileStorage,
	}
}

func (s *FileService) UploadFile(ctx context.Context, file []byte) (string, error) {
	fileID, err := s.fileStorage.UploadFile(ctx, file)
	if err != nil {
		return "", err
	}
	return fileID, nil
}
