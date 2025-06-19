package interfaces

import "context"

type FileStorage interface {
	UploadFile(ctx context.Context, file []byte) (string, error)
}
