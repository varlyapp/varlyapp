package nft

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	r "runtime"

	"github.com/disintegration/imaging"
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

func GenerateNewCollectionFromConfig(ctx context.Context, config NewCollectionConfig) {
	start := time.Now()

	runtime.EventsEmit(ctx, "collection.generation.started", map[string]int{"CollectionSize": config.Size})

	var jobs []Job

	for i := 0; i < config.Size; i++ {
		jobs = append(jobs, Job{Id: i, Config: config})
	}

	var (
		wg              sync.WaitGroup
		jobChannel      = make(chan Job)
		numberOfWorkers = r.NumCPU()
	)

	wg.Add(numberOfWorkers)

	// Start the workers
	for i := 0; i < numberOfWorkers; i++ {
		go nftWorker(i, &wg, jobChannel, ctx)
	}

	// Send jobs to worker
	for _, job := range jobs {
		jobChannel <- job
	}

	close(jobChannel)
	wg.Wait()
	fmt.Printf("Took %s\n", time.Since(start))
}

func ReadLayers(ctx context.Context, dir string) fs.CollectionConfig {
	config, err := fs.ReadLayers(dir)

	if err != nil {
		log.Fatalf("Did not read layers: %s", err)
	}

	return config
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

func makeNFT(workerId int, job Job, ctx context.Context) {
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

	defer func() {
		completed++

		data := map[string]string{"ItemNumber": fmt.Sprint(completed), "CollectionSize": fmt.Sprint(job.Config.Size)}
		runtime.EventsEmit(ctx, "collection.item.generated", data)

		fmt.Printf("%d. %s\n", job.Id, png)
	}()
	GeneratePNG(images, png, job.Config.Width, job.Config.Height)
}

func nftWorker(id int, wg *sync.WaitGroup, jobChannel <-chan Job, ctx context.Context) {
	defer wg.Done()
	for job := range jobChannel {
		makeNFT(id, job, ctx)
	}
}
