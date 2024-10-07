package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadFile(file *multipart.FileHeader) (string, error) {
	uploadPath := "./uploads" 
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(filepath.Join(uploadPath, file.Filename))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return "/uploads/" + file.Filename, nil
}
