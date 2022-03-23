package nft

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/fogleman/gg"
	wr "github.com/mroth/weightedrand"
	"github.com/varlyapp/varlyapp/backend/fs"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type MetaCollection struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Image       string          `json:"image"`
	ExternalURL string          `json:"external_url"`
	Attributes  []MetaAttribute `json:"attributes"`
}

type MetaAttribute struct {
	Type  string `json:"trait_type"`
	Value string `json:"value"`
}

type NewCollectionConfig struct {
	Dir    string
	Order  []string
	Layers map[string][]fs.Layer
	Width  int
	Height int
	Size   int
}

func GenerateCollectionFromConfig(ctx context.Context, config NewCollectionConfig) {
	var waitForPngGeneration sync.WaitGroup

	runtime.EventsEmit(ctx, "collection.generation.started", map[string]int{"CollectionSize": config.Size})
	rand.Seed(time.Now().UTC().UnixNano())

	items := 0
	completed := 0

	for items < config.Size {
		fmt.Printf("%d\n", items)
		var images []string

		for _, trait := range config.Order {
			files := config.Layers[trait]

			if len(files) > 0 {
				var choices []wr.Choice

				for _, layer := range files {
					choices = append(choices, wr.Choice{Item: layer.Item, Weight: uint(layer.Weight)})
				}

				chooser, err := wr.NewChooser(choices...)

				if err != nil {
					log.Fatal(err)
				}

				pick := chooser.Pick().(string)

				images = append(images, pick)
			}
		}

		waitForPngGeneration.Add(1)

		go func(items int) {
			pngfile := fmt.Sprintf("%s/%d.png", config.Dir, items)
			jsonfile := fmt.Sprintf("%s/%d.json", config.Dir, items)

			defer func(file string) {
				completed++
				data := map[string]string{"ItemNumber": fmt.Sprint(completed), "CollectionSize": fmt.Sprint(config.Size), "ImagePath": pngfile}
				runtime.EventsEmit(ctx, "collection.item.generated", data)
				fmt.Printf("Saved %s ☑️\n", file)
				waitForPngGeneration.Done()
			}(pngfile)

			GenerateMetadata(MetaCollection{
				Name:        "My Collection",
				Description: "Description for my collection goes here.",
				Image:       pngfile,
				ExternalURL: fmt.Sprintf("https://mynft.com?token=%d", items),
				Attributes: []MetaAttribute{
					{Type: "Layer 1", Value: "Some Value"},
					{Type: "Layer 2", Value: "Some Value"},
					{Type: "Layer 3", Value: "Some Value"},
					{Type: "Layer 4", Value: "Some Value"},
				},
			}, jsonfile)
			GeneratePNG(images, pngfile, config.Width, config.Height)
		}(items)

		items++
	}

	waitForPngGeneration.Wait()
	fmt.Println("👍")
}

func ReadLayers(ctx context.Context, dir string) fs.CollectionConfig {
	config, err := fs.ReadLayers(dir)

	if err != nil {
		log.Fatalf("Did not read layers: %s", err)
	}

	return config
}

func GenerateMetadata(metadata MetaCollection, output string) bool {
	b, err := json.Marshal(metadata)

	if err != nil {
		return false
	}

	err = os.WriteFile(output, b, os.ModePerm)

	if err != nil {
		return false
	}

	return true
}

func GeneratePNG(layers []string, output string, width int, height int) bool {
	context := gg.NewContext(width, height)

	for _, img := range layers {
		context.DrawImage(decode(img), 0, 0)
	}

	err := context.SavePNG(output)

	if err != nil {
		log.Fatalf("Unable to save image %s", err)
		return false
	}

	return true
}

func decode(path string) image.Image {
	img, err := os.Open(path)

	if err != nil {
		log.Fatalf("Failed to open %s", err)
	}

	defer img.Close()

	decoded, err := png.Decode(img)

	if err != nil {
		log.Fatalf("Failed to decode %s", err)
	}

	return decoded
}
