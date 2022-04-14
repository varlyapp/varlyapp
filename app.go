package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/varlyapp/varlyapp/backend/services"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	userdir, _ = os.UserConfigDir()
	basedir    = filepath.Join(userdir, "varlyapp")
	docsdir    = filepath.Join(basedir, "Documents")
)

// App struct
type App struct {
	ctx               context.Context
	Docs              map[string]interface{}
	SettingsService   *services.SettingsService
	FileSystemService *services.FileSystemService
	CollectionService *services.CollectionService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		SettingsService:   services.NewSettingsService(docsdir),
		FileSystemService: services.NewFileSystemService(),
		CollectionService: services.NewCollectionService(),
	}
}

// startup is called at application startup
func (app *App) startup(ctx context.Context) {
	// Perform your setup here
	app.ctx = ctx
	app.FileSystemService.Ctx = ctx
	app.CollectionService.Ctx = ctx
	m := menu.NewMenuFromItems(
		menu.AppMenu(),
		menu.SubMenu("File", menu.NewMenuFromItems(
			menu.Text("Open Collection", keys.CmdOrCtrl("o"), func(cd *menu.CallbackData) {
				runtime.EventsEmit(ctx, "shortcut.collection.open")
			}),
			menu.Text("Save Collection", keys.CmdOrCtrl("s"), func(cd *menu.CallbackData) {
				runtime.EventsEmit(ctx, "shortcut.collection.save")
			}),
		)),
		menu.EditMenu(),
	)

	runtime.MenuSetApplicationMenu(ctx, m)
}

// domReady is called after the front-end dom has been loaded
func (app *App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (app *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (app *App) OpenDirectoryDialog(title string) string {
	path, _ := runtime.OpenDirectoryDialog(app.ctx, runtime.OpenDialogOptions{
		Title:                      title,
		CanCreateDirectories:       true,
		TreatPackagesAsDirectories: true,
	})

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

func (app *App) EncodeImage(path string) string {
	image, err := os.ReadFile(path)

	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}

	encoded := base64.StdEncoding.EncodeToString(image)

	encoded = fmt.Sprintf("data:image/png;base64,%s", encoded)

	return encoded
}

func (app *App) SaveFile(file string, data string) bool {
	path := fmt.Sprintf("%s%s", docsdir, file)

	err := os.WriteFile(path, []byte(data), os.ModePerm)

	if err != nil {
		return false
	}

	return true
}

func (app *App) GetImageStats(path string) image.Config {
	reader, err := os.Open(path)
	if err != nil {
		return image.Config{}
	}
	defer reader.Close()
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		return image.Config{}
	}

	return img
}

func (app *App) MessageDialog(options runtime.MessageDialogOptions) string {
	res, _ := runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         options.Title,
		Message:       options.Message,
		Buttons:       options.Buttons,
		DefaultButton: options.DefaultButton,
	})

	return res
}
