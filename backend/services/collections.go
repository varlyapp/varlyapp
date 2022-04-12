package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"

	wr "github.com/mroth/weightedrand"
	"github.com/pkg/errors"
	"github.com/varlyapp/varlyapp/backend/lib"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type VariantFileType string

const (
	VariantFileTypePng  VariantFileType = "PNG"
	VariantFileTypeTiff                 = "TIFF"
	VariantFileTypeGif                  = "GIF"
)

const FileTypeExpression = `.(png|jpg|jpeg)$`
const RarityExpression = `(.*)#([.0-9]+).(png|jpg|jpeg)$`
const DefaultRarity float64 = 100.0

type CollectionService struct {
	Ctx        context.Context
	Collection *Collection
}

type Trait struct {
	Name      string `json:"name"`
	Collapsed bool   `json:"collapsed"`
}

type Variant struct {
	Name   string  `json:"name"`
	Path   string  `json:"path"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
}

type Collection struct {
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	SourceDirectory string               `json:"sourceDirectory"`
	OutputDirectory string               `json:"outputDirectory"`
	Traits          []Trait              `json:"traits"`
	Layers          map[string][]Variant `json:"layers"`
	LayersDummy     []Variant            `json:"layersDummy"`
	Width           float64              `json:"width"`
	Height          float64              `json:"height"`
	Size            int                  `json:"size"`
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func NewCollectionService() *CollectionService {
	return &CollectionService{}
}

func GetVariantDefaults() *Variant {
	return &Variant{
		Name:   "",
		Path:   "",
		Width:  0.0,
		Height: 0.0,
		// FileType:   VariantFileTypePng,
	}
}

func GetCollectionDefaults() *Collection {
	return &Collection{
		SourceDirectory: "",
		OutputDirectory: "",
		Traits:          []Trait{},
		Layers:          map[string][]Variant{},
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

// ReadsDirectory reads a direcotory into a Collection
// Items within the directory are split into Layers
func (c *CollectionService) LoadCollectionFromDirectory(dir string) *Collection {
	collection := &Collection{Layers: map[string][]Variant{}}
	lastDirectory := ""

	fileX := regexp.MustCompile(FileTypeExpression)
	rarityX := regexp.MustCompile(RarityExpression)

	err := filepath.Walk(
		dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				if dir != path {
					lastDirectory = info.Name()
				}
			} else {
				if fileX.MatchString(info.Name()) {
					var weight float64

					if rarityX.MatchString(path) {
						weight, _ = strconv.ParseFloat(rarityX.ReplaceAllString(path, "$2"), 64)
					}

					if weight <= 0.0 {
						weight = DefaultRarity
					}

					collection.Layers[lastDirectory] = append(collection.Layers[lastDirectory], Variant{Name: info.Name(), Path: path, Weight: weight})
				}
			}

			return nil
		})

	if err != nil {
		lib.ShowErrorModal(c.Ctx, "Collection cannot be loaded from directory", err.Error())
		return nil
	}

	return collection
}

func (c *CollectionService) SaveCollection(collection *Collection) error {
	path, err := runtime.SaveFileDialog(c.Ctx, runtime.SaveDialogOptions{
		Title:           "Save Varly collection as a file",
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

	contents, err := json.MarshalIndent(collection, "", "  ")

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

func (c *CollectionService) GenerateCollection(collection Collection) {
	runtime.EventsEmit(c.Ctx, "collection.generation.started", map[string]int{"CollectionSize": collection.Size})

	var jobs []lib.Job

	for i := 0; i < collection.Size; i++ {
		jobs = append(jobs, lib.Job{Id: i, Config: collection})
	}

	var completed = 0
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		lib.Batch(c.Ctx, 0, jobs, func(ctx context.Context, id int, job lib.Job) {
			var images []string
			collection := job.Config.(Collection)

			for _, trait := range collection.Traits {
				variants := collection.Layers[trait.Name]

				if len(variants) > 0 {
					var choices []wr.Choice

					for _, variant := range variants {
						choices = append(choices, wr.Choice{Item: variant.Path, Weight: uint(variant.Weight)})
					}

					chooser, err := wr.NewChooser(choices...)

					if err != nil {
						log.Fatal(err)
					}

					pick := chooser.Pick().(string)

					images = append(images, pick)
				}
			}

			png := fmt.Sprintf("%s/%d.png", collection.OutputDirectory, job.Id)

			defer func() {
				completed++

				data := map[string]string{"ItemNumber": fmt.Sprint(completed), "CollectionSize": fmt.Sprint(collection.Size)}
				runtime.EventsEmit(ctx, "collection.item.generated", data)

				fmt.Printf("%d. %s\n", job.Id, png)
			}()

			runtime.EventsEmit(ctx, "debug", map[string]interface{}{
				"images": images, "png": png,
			})
			err := lib.GeneratePNG(images, png, int(collection.Width), int(collection.Height))

			if err != nil {
				fmt.Printf("unable to generate image: %s\n%v\n", png, err)
			}
		})
	}()

	wg.Wait()
}

func (c *CollectionService) GenerateCollectionPreview(collection Collection) string {
	var layers []string

	for _, trait := range collection.Traits {
		variants := collection.Layers[trait.Name]

		if len(variants) > 0 {
			var choices []wr.Choice

			for _, variant := range variants {
				choices = append(choices, wr.Choice{Item: variant.Path, Weight: uint(variant.Weight)})
			}

			chooser, err := wr.NewChooser(choices...)

			if err != nil {
				log.Fatal(err)
			}

			pick := chooser.Pick().(string)

			layers = append(layers, pick)
		}
	}

	preview, err := lib.MakePreview(layers, int(collection.Width), int(collection.Height), 512)

	if err != nil {
		lib.ShowErrorModal(c.Ctx, "No Preview", "Sorry, I was unable to generate preview for this collection")
	}
	return preview
}
