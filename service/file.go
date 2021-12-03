package service

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type File struct{}

func (f File) saveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func (f File) Store(file *multipart.FileHeader, dir string) (string, error) {
	ext := filepath.Ext(file.Filename)
	dst := filepath.Join("uploads", dir, uuid.NewString()+ext)
	err := f.saveUploadedFile(file, dst)

	return dst, err
}
