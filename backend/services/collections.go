package services

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/varlyapp/varlyapp/backend/lib"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type LayerType string

const (
	LayerTypePng  LayerType = "PNG"
	LayerTypeTiff           = "TIFF"
	LayerTypeGif            = "GIF"
)

type Layer struct {
	Name   string    `json:"name"`
	Path   string    `json:"path"`
	Width  float64   `json:"width"`
	Height float64   `json:"height"`
	Weight float64   `json:"weight"`
	Type   LayerType `json:"type"`
}

type Trait struct {
	Name      string `json:"name"`
	Collapsed bool   `json:"collapsed"`
}

type Variant struct{}

type Collection struct {
	Name            string             `json:"name"`
	Description     string             `json:"description"`
	SourceDirectory string             `json:"sourceDirectory"`
	OutputDirectory string             `json:"outputDirectory"`
	Traits          []Trait            `json:"traits"`
	Layers          map[string][]Layer `json:"layers"`
	Width           float64            `json:"width"`
	Height          float64            `json:"height"`
	Size            int                `json:"size"`
}

type CollectionService struct {
	Ctx        context.Context
	Collection *Collection
}

func NewCollectionService() *CollectionService {
	return &CollectionService{}
}

func GetLayerDefaults() *Layer {
	return &Layer{
		Name:   "",
		Path:   "",
		Width:  0.0,
		Height: 0.0,
		Type:   LayerTypePng,
	}
}
func GetCollectionDefaults() *Collection {
	return &Collection{
		SourceDirectory: "",
		OutputDirectory: "",
		Traits:          []Trait{},
		Layers:          make(map[string][]Layer, 0),
		Width:           0.0,
		Height:          0.0,
		Size:            100,
	}
}
func (c *CollectionService) LoadCollection() *Collection {
	collection := &Collection{}
	content, err := lib.OpenFileContents(c.Ctx)

	if err != nil {
		lib.ShowErrorModal(c.Ctx, "Collection cannot be loaded", "The collection file may be corrupted or data may be missing")
		return nil
	} else {
		err = json.Unmarshal([]byte(content), collection)

		if err != nil {
			lib.ShowErrorModal(c.Ctx, "Collection cannot be read", err.Error())
			return nil
		}

		return collection
	}
}
func (c *CollectionService) SaveCollection(collection *Collection) error {
	path, err := runtime.SaveFileDialog(c.Ctx, runtime.SaveDialogOptions{
		Title: "Save Varly collection as a file",
		DefaultFilename: collection.Name,
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Varly Collection Files (*.json, *.varly)",
				Pattern:     "*.json;*.varly",
			},
		},
	})
	if err != nil {
		return err
	}
	if path == "" {
		return errors.New("path to file was empty")
	}

	contents, err := json.Marshal(collection)

	if err != nil {
		return err
	}

	err = lib.WriteFileContents(path, contents)

	if err != nil {
		lib.ShowErrorModal(c.Ctx, "Collection cannot be saved", "Collection data may be corrupted")
		return nil
	}

	return nil
}
func (c *CollectionService) ValidateCollection() error {
	return nil
}
