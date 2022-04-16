package services

import (
	"context"
	"fmt"

	"github.com/varlyapp/varlyapp/backend/lib"
)

type FileSystemService struct {
	Ctx context.Context
}

func NewFileSystemService() *FileSystemService {
	return &FileSystemService{}
}

func (f *FileSystemService) ReadFileContents(path string) string {
	contents, err := lib.ReadFileContents(path)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return contents
}

func (f *FileSystemService) OpenFileContents() string {
	contents, _ := lib.OpenFileContents(f.Ctx)

	return contents
}
