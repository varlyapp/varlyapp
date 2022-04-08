package lib

import (
	"context"
	"errors"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func ReadFileContents(path string) (string, error) {
	contents, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(contents), nil
}

func OpenFileContents(ctx context.Context) (string, error) {
	path, err := runtime.OpenFileDialog(ctx, runtime.OpenDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Varly Collection Files (*.json, *.varly)",
				Pattern:     "*.json;*.varly",
			},
		},
	})
	if err != nil {
		return "", err
	}
	if path == "" {
		return "", errors.New("path to file was empty")
	}
	contents, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(contents), nil
}
