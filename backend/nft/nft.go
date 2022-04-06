package nft

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"os"
	r "runtime"
	"strconv"
	"sync"
	"time"

	"github.com/disintegration/imaging"
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

type Job struct {
	Id     int
	Config NewCollectionConfig
}

type JobResult struct {
	Output string
}

var (
	completed = 0
	// numberOfWorkers = 10
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func GenerateCollectionFromConfig(ctx context.Context, config NewCollectionConfig) {
	runtime.EventsEmit(ctx, "collection.generation.started", map[string]int{"CollectionSize": config.Size})

	pool := queue.New(r.NumCPU())

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go pool.GenerateFrom(tasks(config))
	go pool.Run(ctx)

	for {
		select {
		case r, open := <-pool.Results():
			if !open {
				continue
			}

			i, err := strconv.ParseInt(string(r.Task.ID), 10, 64)
			if err != nil {
				log.Fatalf("unexpected integer parsing error: %v", err)
			}

			val := r.Value.(int)
			if val != int(i) {
				log.Fatalf("wrong value %v; expected %v", val, int(i))
			}
		case <-pool.Done:
			return
		default:
		}
	}
}

func GenerateNewCollectionFromConfig(ctx context.Context, config NewCollectionConfig) {
	runtime.EventsEmit(ctx, "collection.generation.started", map[string]int{"CollectionSize": config.Size})

	start := time.Now()
	// Create fake jobs for testing purposes
	var jobs []Job

	// for i := 0; i < 1000; i++ {
	// 	jobs = append(jobs, Job{Id: i})
	// }

	for i := 0; i < config.Size; i++ {
		jobs = append(jobs, Job{Id: i, Config: config})
	}

	var (
		wg              sync.WaitGroup
		jobChannel      = make(chan Job)
		numberOfWorkers = 16
	)

	runtime.LogInfo(ctx, fmt.Sprintf("%v", numberOfWorkers))
	wg.Add(numberOfWorkers)

	// Start the workers
	for i := 0; i < numberOfWorkers; i++ {
		go worker(i, &wg, jobChannel, ctx)
	}

	// Send jobs to worker
	for _, job := range jobs {
		jobChannel <- job
	}

	close(jobChannel)
	wg.Wait()
	fmt.Printf("Took %s\n", time.Since(start))
}

func worker(id int, wg *sync.WaitGroup, jobChannel <-chan Job, ctx context.Context) {
	defer wg.Done()
	for job := range jobChannel {
		generateNewArtwork(id, job, ctx)
	}
}

func generateNewArtwork(workerId int, job Job, ctx context.Context) {
	var images []string

	for _, trait := range job.Config.Order {
		files := job.Config.Layers[trait]

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

	png := fmt.Sprintf("%s/%d.png", job.Config.Dir, job.Id)
	// fmt.Printf("Worker #%d Running job #%d\n", workerId, job.Id)
	// time.Sleep(500 * time.Millisecond)
	// return JobResult{Output: "Success"}
	defer func() {
		completed++

		data := map[string]string{"ItemNumber": fmt.Sprint(completed), "CollectionSize": fmt.Sprint(job.Config.Size)}
		runtime.EventsEmit(ctx, "collection.item.generated", data)

		fmt.Printf("%d. %s\n", job.Id, png)
	}()
	GeneratePNG(images, png, job.Config.Width, job.Config.Height)
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
	png := imaging.New(width, height, color.NRGBA{0, 0, 0, 0})

	for _, layer := range layers {
		img, _ := imaging.Open(layer)
		png = imaging.Overlay(png, img, image.Pt(0, 0), 1)
	}

	err := imaging.Save(png, output)
	if err != nil {
		fmt.Println("unable to save image: ", err.Error())
		return false
	}
	return true
}

func tasks(config NewCollectionConfig) []queue.Task {
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
		return nil, errors.New("unexpected error fetching arguments from task")
	}

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
			completed++

			data := map[string]string{"ItemNumber": fmt.Sprint(completed), "CollectionSize": fmt.Sprint(t.Metadata.Size)}
			runtime.EventsEmit(ctx, "collection.item.generated", data)

			fmt.Printf("%d. %s\n", val, png)
			wg.Done()
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

	return val, nil
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
