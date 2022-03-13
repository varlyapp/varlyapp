package main

import (
	"context"

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
	nft.GenerateCollectionFromConfig(app.ctx, config);
}

func (app *App) GenerateCollection(dir string) fs.CollectionConfig {
	return nft.GenerateCollection(app.ctx, dir, []string{
		"backgrounds",
		"shapes",
		// "Background",
		// "Eyeball",
		// "Eye color",
		// "Iris",
		// "Shine",
		// "Bottom lid",
		// "Top lid",
	}, 1500, 1500, 10);
}

func (app *App) UploadCollection() string {
	return fs.UploadCollection();
}
