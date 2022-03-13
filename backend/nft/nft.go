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
	var wg sync.WaitGroup

	runtime.EventsEmit(ctx, "collection.generation.started", map[string]int{"CollectionSize": config.Size})

	// config, err := fs.ReadLayers(dir)

	// if err != nil {
	// 	log.Fatalf("Did not read layers: %s", err)
	// }

	// return config

	rand.Seed(time.Now().UTC().UnixNano())

	items := 0
	completed := 0

	for items < config.Size {
		wg.Add(1)

		var images []string

		for _, trait := range config.Order {
			fmt.Printf("Processing layer: %s\n\n", trait)
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

		fmt.Println("Should have finished setting up images")

		go func(i int) {
			defer func() {
				completed += 1

				data := map[string]int{"ItemNumber": completed, "CollectionSize": config.Size}
				runtime.EventsEmit(ctx, "collection.item.generated", data)
				wg.Done()
			}()

			output := fmt.Sprintf("%s/%d.png", config.Dir, i)
			GeneratePNG(images, output, config.Width, config.Height)
		}(items)

		items++
	}

	wg.Wait()
	fmt.Println("👍")
}

func GenerateCollection(ctx context.Context, dir string, order []string, width int, height int, size int) fs.CollectionConfig {
	var wg sync.WaitGroup

	runtime.EventsEmit(ctx, "collection.generation.started", map[string]int{"CollectionSize": size})

	config, err := fs.ReadLayers(dir)

	if err != nil {
		log.Fatalf("Did not read layers: %s", err)
	}

	return config

	rand.Seed(time.Now().UTC().UnixNano())

	items := 0
	completed := 0

	for items < size {
		wg.Add(1)

		var images []string

		for _, trait := range order {
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

		go func(i int) {
			defer func() {
				completed += 1

				data := map[string]int{"ItemNumber": completed, "CollectionSize": size}
				runtime.EventsEmit(ctx, "collection.item.generated", data)
				wg.Done()
			}()

			output := fmt.Sprintf("./output/%d.png", i)
			GeneratePNG(images, output, width, height)
		}(items)

		items++
	}

	wg.Wait()
	fmt.Println("👍")

	return config
}

func GeneratePNG(layers []string, output string, width int, height int) {
	dc := gg.NewContext(width, height)

	for _, img := range layers {
		fmt.Printf("Drawing image: %s\n", img)
		dc.DrawImage(decode(img), 0, 0)
	}

	err := dc.SavePNG(output)

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
