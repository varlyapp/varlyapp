package lib

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func GetApplicationDocumentsDirectory(paths ...string) (string, error) {
	dir, _ := os.UserConfigDir()

	path := fmt.Sprintf("%s/varlyapp/Documents/%s", dir, strings.Join(paths, "/"))

	err := os.MkdirAll(path, os.ModePerm);

	if err != nil {
		return "", err
	}

	return path, nil
}
func ReadFileContents(path string) (string, error) {
	contents, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(contents), nil
}

func WriteFileContents(path string, contents []byte) error {
	err := os.WriteFile(path, contents, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
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
