package nft

import (
	"context"
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
			output := fmt.Sprintf("%s/%d.png", config.Dir, items)

			defer func(file string) {
				completed++
				data := map[string]int{"ItemNumber": completed, "CollectionSize": config.Size}
				runtime.EventsEmit(ctx, "collection.item.generated", data)
				fmt.Printf("Saved %s ☑️\n", file)
				waitForPngGeneration.Done()
			}(output)

			GeneratePNG(images, output, config.Width, config.Height)
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

func GeneratePNG(layers []string, output string, width int, height int) {
	context := gg.NewContext(width, height)

	for _, img := range layers {
		context.DrawImage(decode(img), 0, 0)
	}

	err := context.SavePNG(output)

	if err != nil {
		log.Fatalf("Unable to save image %s", err)
	}
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
