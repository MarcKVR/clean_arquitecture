package repositories

import "mime/multipart"

type FileRepository interface {
	UploadFile(file *multipart.FileHeader, destination string) (string, error)
}
