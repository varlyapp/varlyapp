package nft

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	r "runtime"
	"strconv"
	"sync"
	"time"

	"github.com/davidbyttow/govips/v2/vips"
	wr "github.com/mroth/weightedrand"
	"github.com/varlyapp/varlyapp/backend/fs"
	"github.com/varlyapp/varlyapp/backend/queue"
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

var completed = 0

func GenerateCollectionFromConfig(ctx context.Context, config NewCollectionConfig) {
	vips.Startup(&vips.Config{
		ReportLeaks: true,
		CollectStats: true,
	})
	defer vips.Shutdown()

	// Set CPU max to half or 1, whichever is greater
	// cpu := r.NumCPU()
	// max := math.Max(float64(1), math.Floor(float64(cpu) / 2))
	// r.GOMAXPROCS(int(max))
	rand.Seed(time.Now().UTC().UnixNano())

	runtime.EventsEmit(ctx, "collection.generation.started", map[string]int{"CollectionSize": config.Size})

	pool := queue.New(5)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go pool.GenerateFrom(testTasks(config))
	go pool.Run(ctx)

	for {
		select {
		case r, open := <-pool.Results():
			if !open {
				continue
			}

			i, err := strconv.ParseInt(string(r.Task.ID), 10, 64)
			if err != nil {
				log.Fatalf("unexpected error: %v", err)
			}

			val := r.Value.(int)
			if val != int(i)*2 {
				log.Fatalf("wrong value %v; expected %v", val, int(i)*2)
			}
		case <-pool.Done:
			fmt.Println("Done 👍")

			return
		default:
		}
	}
}

func GenerateCollectionFromConfig2(ctx context.Context, config NewCollectionConfig) {
	// Set CPU max to half or 1, whichever is greater
	cpu := r.NumCPU()
	max := math.Max(float64(1), math.Floor(float64(cpu)/2))
	r.GOMAXPROCS(int(max))

	var wg sync.WaitGroup

	runtime.EventsEmit(ctx, "collection.generation.started", map[string]int{"CollectionSize": config.Size})
	rand.Seed(time.Now().UTC().UnixNano())

	items := 0
	completed := 0

	for items < config.Size {
		wg.Add(1)

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

		go func(items int) {
			pngfile := fmt.Sprintf("%s/%d.png", config.Dir, items)
			// jsonfile := fmt.Sprintf("%s/%d.json", config.Dir, items)

			defer func(file string) {
				completed++
				data := map[string]string{"ItemNumber": fmt.Sprint(completed), "CollectionSize": fmt.Sprint(config.Size), "ImagePath": pngfile}
				runtime.EventsEmit(ctx, "collection.item.generated", data)
				fmt.Printf("Saved %d ☑️\n", completed)
				wg.Done()
			}(pngfile)

			// GenerateMetadata(MetaCollection{
			// 	Name:        "My Collection",
			// 	Description: "Description for my collection goes here.",
			// 	Image:       pngfile,
			// 	ExternalURL: fmt.Sprintf("https://mynft.com?token=%d", items),
			// 	Attributes: []MetaAttribute{
			// 		{Type: "Layer 1", Value: "Some Value"},
			// 		{Type: "Layer 2", Value: "Some Value"},
			// 		{Type: "Layer 3", Value: "Some Value"},
			// 		{Type: "Layer 4", Value: "Some Value"},
			// 	},
			// }, jsonfile)
			GeneratePNG(images, pngfile, config.Width, config.Height)
		}(items)

		items++
	}

	wg.Wait()
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
	img, _ := vips.NewImageFromFile(layers[0])

	for _, layer := range layers[1:] {
		i, _ := vips.LoadImageFromFile(layer, vips.NewImportParams())
		img.Composite(i, vips.BlendModeOver, 0, 0)
	}

	bytes, _, err := img.ExportPng(vips.NewPngExportParams())

	if err != nil {
		log.Fatalf("Unable to save image %s", err)
		return false
	}

	os.WriteFile(output, bytes, os.ModePerm)

	return true
}

func testTasks(config NewCollectionConfig) []queue.Task {
	tasks := make([]queue.Task, config.Size)
	for i := 0; i < config.Size; i++ {
		tasks[i] = queue.Task{
			ID:   queue.TaskID(fmt.Sprintf("%v", i)),
			Type: "anyType",
			Metadata: queue.TaskMetadata{
				Dir:    config.Dir,
				Order:  config.Order,
				Layers: config.Layers,
				Width:  config.Width,
				Height: config.Height,
				Size:   config.Size,
			},
			Callback:  callback,
			Arguments: i,
		}
	}
	return tasks
}

func callback(ctx context.Context, t queue.Task) (interface{}, error) {
	val, ok := t.Arguments.(int)
	if !ok {
		return nil, errors.New("Something didn't go right")
	}

	// time.Sleep(250 * time.Millisecond)
	// fmt.Printf("%d. /Documents/dir/%d.png ☑️\n", val, val)
	// return val * 2, nil

	var images []string

	for _, trait := range t.Metadata.Order {
		files := t.Metadata.Layers[trait]

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

	png := fmt.Sprintf("%s/%d.png", t.Metadata.Dir, val)
	// json := fmt.Sprintf("%s/%d.json", t.Metadata.Dir, val)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer func() {
			wg.Done()

			completed++

			data := map[string]string{"ItemNumber": fmt.Sprint(completed), "CollectionSize": fmt.Sprint(t.Metadata.Size)}
			runtime.EventsEmit(ctx, "collection.item.generated", data)

			fmt.Printf("%d. %s\n", val, png)
		}()
		GeneratePNG(images, png, t.Metadata.Width, t.Metadata.Height)
	}()

	// go func(j string, p string) {
	// 	defer func() {
	// 		wg.Done()
	// 		fmt.Printf("%d. %s\n", val, json)
	// 	}()

	// 	meta := MetaCollection{
	// 		Name:        "My Collection",
	// 		Description: "Description for my collection goes here.",
	// 		Image:       p,
	// 		ExternalURL: fmt.Sprintf("https://mynft.com?token=%d", val),
	// 		Attributes: []MetaAttribute{
	// 			{Type: "Layer 1", Value: "Some Value"},
	// 			{Type: "Layer 2", Value: "Some Value"},
	// 			{Type: "Layer 3", Value: "Some Value"},
	// 			{Type: "Layer 4", Value: "Some Value"},
	// 		},
	// 	}

	// 	GenerateMetadata(meta, j)
	// }(json, png)

	wg.Wait()

	return val * 2, nil
	// go func(items int) {
	// 	pngfile := fmt.Sprintf("%s/%d.png", config.Dir, items)
	// jsonfile := fmt.Sprintf("%s/%d.json", config.Dir, items)

	// defer func(file string) {
	// 	completed++
	// 	data := map[string]string{"ItemNumber": fmt.Sprint(completed), "CollectionSize": fmt.Sprint(config.Size), "ImagePath": pngfile}
	// 	runtime.EventsEmit(ctx, "collection.item.generated", data)
	// 	fmt.Printf("Saved %d ☑️\n", completed)
	// 	wg.Done()
	// }(pngfile)

	// GenerateMetadata(MetaCollection{
	// 	Name:        "My Collection",
	// 	Description: "Description for my collection goes here.",
	// 	Image:       pngfile,
	// 	ExternalURL: fmt.Sprintf("https://mynft.com?token=%d", items),
	// 	Attributes: []MetaAttribute{
	// 		{Type: "Layer 1", Value: "Some Value"},
	// 		{Type: "Layer 2", Value: "Some Value"},
	// 		{Type: "Layer 3", Value: "Some Value"},
	// 		{Type: "Layer 4", Value: "Some Value"},
	// 	},
	// }, jsonfile)
	// GeneratePNG(images, pngfile, config.Width, config.Height)
	// }(items)
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
