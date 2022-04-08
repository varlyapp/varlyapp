package services

import (
	"context"
	"encoding/json"

	"github.com/varlyapp/varlyapp/backend/lib"
)

type LayerType string

const (
	LayerTypePng  LayerType = "PNG"
	LayerTypeTiff           = "TIFF"
	LayerTypeGif            = "GIF"
)

type Layer struct {
	Name   string
	Path   string
	Width  float64
	Height float64
	Weight float64
	Type   LayerType
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
	Width           int                `json:"width"`
	Height          int                `json:"height"`
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
	return nil
}
func (c *CollectionService) ValidateCollection() error {
	return nil
}
