package services

import (
	"context"

	"github.com/varlyapp/varlyapp/backend/lib"
)

type FileSystemService struct {
	Ctx context.Context
}

func NewFileSystemService() *FileSystemService {
	return &FileSystemService{}
}

func (f *FileSystemService) ReadFileContents(path string) string {
	contents, _ := lib.ReadFileContents(path)

	return contents
}

func (f *FileSystemService) OpenFileContents() string {
	contents, _ := lib.OpenFileContents(f.Ctx)

	return contents
}
