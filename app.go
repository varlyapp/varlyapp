package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/varlyapp/varlyapp/backend/fs"
	"github.com/varlyapp/varlyapp/backend/nft"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (app *App) startup(ctx context.Context) {
	// Perform your setup here
	app.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (app *App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (app *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (app *App) OpenDirectoryDialog() string {
	path, _ := runtime.OpenDirectoryDialog(app.ctx, runtime.OpenDialogOptions{})

	return path
}

func (app *App) OpenFileDialog() string {
	path, _ := runtime.OpenFileDialog(app.ctx, runtime.OpenDialogOptions{})

	return path
}

func (app *App) SaveFileDialog() string {
	path, _ := runtime.SaveFileDialog(app.ctx, runtime.SaveDialogOptions{})

	return path
}

func (app *App) GenerateCollectionFromConfig(config nft.NewCollectionConfig) {
	nft.GenerateCollectionFromConfig(app.ctx, config)
}

func (app *App) ReadLayers(dir string) fs.CollectionConfig {
	return nft.ReadLayers(app.ctx, dir)
}

func (app *App) UploadCollection() string {
	return fs.UploadCollection()
}

func (app *App) GetApplicationDocumentsDirectory() string {
	return fs.GetApplicationDocumentsDirectory()
}

func (app *App) EncodeImage(path string) string {
	image, err := ioutil.ReadFile(path)

	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}

	encoded := base64.StdEncoding.EncodeToString(image)

	encoded = fmt.Sprintf("data:image/png;base64,%s", encoded)

	return encoded
}
