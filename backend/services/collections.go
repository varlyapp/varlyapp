package services

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/disintegration/imaging"
	wr "github.com/mroth/weightedrand"
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
const ImageTypeExpression = `.(png|jpg|jpeg)$`
const JsonTypeExpression = `.(json)$`
const RarityExpression = `(.*)#([.0-9]+).(png|jpg|jpeg)$`
const DefaultRarity float64 = 100.0

type CollectionService struct {
	Ctx           context.Context
	Collection    *Collection
	DocsDirecotry string
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
	Artist          string               `json:"artist"`
	BaseUri         string               `json:"baseUri"`
	SourceDirectory string               `json:"sourceDirectory"`
	OutputDirectory string               `json:"outputDirectory"`
	Traits          []Trait              `json:"traits"`
	Layers          map[string][]Variant `json:"layers"`
	LayersDummy     []Variant            `json:"layersDummy"`
	Width           float64              `json:"width"`
	Height          float64              `json:"height"`
	Size            int                  `json:"size"`
}

type FileItem struct {
	filepath string
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func NewCollectionService(docsdir string) *CollectionService {
	return &CollectionService{DocsDirecotry: docsdir}
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
		lib.ErrorModal(c.Ctx, "Collection cannot be loaded", "The collection file may be corrupted or data may be missing")
		return nil
	} else {
		err = json.Unmarshal([]byte(content), collection)

		if err != nil {
			lib.ErrorModal(c.Ctx, "Collection cannot be read", err.Error())
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
		lib.ErrorModal(c.Ctx, "Collection cannot be loaded from directory", err.Error())
		return nil
	}

	return collection
}

func (c *CollectionService) SaveCollection(collection *Collection) string {
	file, err := runtime.SaveFileDialog(c.Ctx, runtime.SaveDialogOptions{
		Title:           "Save Varly collection as a file",
		DefaultFilename: collection.Name,
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Varly Collection Files (*.json, *.varly)",
				Pattern:     "*.json;*.varly",
			},
		},
	})
	fmt.Println(file)
	if err != nil {
		return ""
	}
	if file == "" {
		lib.ErrorModal(c.Ctx, "Collection has no save location", "Collection destination may be read-only")

		return ""
	}

	contents, err := json.MarshalIndent(collection, "", "  ")

	if err != nil {
		lib.ErrorModal(c.Ctx, "Collection could not be formatted", "Collection data may be in a different format")
		return ""
	}

	err = lib.WriteFileContents(file, contents)

	if err != nil {
		lib.ErrorModal(c.Ctx, "Collection could not be saved", "Collection data may be corrupted")
		return ""
	}

	return file
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

	var (
		wg        sync.WaitGroup
		usedDNA   = sync.Map{}
		completed = 0
	)

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

			defer func() {
				completed++
				data := map[string]string{"ItemNumber": fmt.Sprint(completed), "CollectionSize": fmt.Sprint(collection.Size)}
				runtime.EventsEmit(ctx, "collection.item.generated", data)
			}()

			pngFilepath := fmt.Sprintf("%s/%d.png", collection.OutputDirectory, job.Id)

			dna := lib.GenerateDNA(images)
			val, exists := usedDNA.Load(dna)
			if exists {
				pngFilepath = fmt.Sprintf("%s/duplicate-%d.png", collection.OutputDirectory, job.Id)

				fmt.Println("DNA already exists: ", val)
			} else {
				usedDNA.Store(dna, dna)
			}

			runtime.EventsEmit(ctx, "debug", map[string]interface{}{
				"images": images, "png": pngFilepath,
			})

			err1 := lib.GenerateMetadata(lib.MetadataConfig{
				Name:        collection.Name,
				Description: collection.Description,
				BaseURI:     collection.BaseUri,
				Edition:     job.Id,
				Layers:      images,
				Image:       pngFilepath,
				Artist:      collection.Artist,
				DNA:         dna,
			})

			if err1 != nil {
				fmt.Printf("unable to generate metadata: %s\n%v\n", pngFilepath, err1)
			}

			err2 := lib.GeneratePNG(images, pngFilepath, int(collection.Width), int(collection.Height))

			if err2 != nil {
				fmt.Printf("unable to generate image: %s\n%v\n", pngFilepath, err2)
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
		fmt.Println(err)
		lib.ErrorModal(c.Ctx, "No Preview", "Could not generate preview")
	}
	return preview
}

func (c *CollectionService) GenerateCollectionGif(collection Collection, fps int, delay int) string {
	// layers := []string{
	// 	"/Users/selvinortiz/Desktop/hashlips output/0.png",
	// 	"/Users/selvinortiz/Desktop/hashlips output/1.png",
	// 	"/Users/selvinortiz/Desktop/hashlips output/3.png",
	// 	"/Users/selvinortiz/Desktop/hashlips output/4.png",
	// 	"/Users/selvinortiz/Desktop/hashlips output/5.png",
	// 	"/Users/selvinortiz/Desktop/hashlips output/6.png",
	// 	"/Users/selvinortiz/Desktop/hashlips output/8.png",
	// 	"/Users/selvinortiz/Desktop/hashlips output/9.png",
	// 	"/Users/selvinortiz/Desktop/hashlips output/11.png",
	// 	"/Users/selvinortiz/Desktop/hashlips output/12.png",
	// }

	runtime.LogInfo(c.Ctx, "Starting job")

	var wg sync.WaitGroup
	var images []string

	wg.Add(fps)
	for i := 0; i < fps; i++ {
		runtime.LogInfo(c.Ctx, "Generating images")

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

		path := filepath.Join(os.TempDir(), fmt.Sprintf("%d.png", i))

		go func() {
			defer wg.Done()
			if err := lib.GenerateFrame(layers, path, int(collection.Width), int(collection.Height), 512); err != nil {
				panic(err)
			}

			images = append(images, path)
			runtime.LogInfo(c.Ctx, "Generated image")
		}()
	}

	wg.Wait()
	wg.Add(len(images))

	var frames []*image.Paletted
	for _, img := range images {
		runtime.LogInfo(c.Ctx, "Generating frames")

		go func(i string) {
			defer wg.Done()

			img, err := imaging.Open(i)
			if err != nil {
				log.Printf("Skipping file %s due to error reading it :%s", i, err)
			}

			buf := bytes.Buffer{}
			if err := gif.Encode(&buf, img, nil); err != nil {
				log.Printf("Skipping file %s due to error in gif encoding:%s", i, err)
			}
			decoded, err := gif.Decode(&buf)
			if err != nil {
				log.Printf("Skipping file %s due to weird error reading the temporary gif :%s", i, err)
			}
			frames = append(frames, decoded.(*image.Paletted))
			runtime.LogInfo(c.Ctx, "Generated frame")
		}(img)
	}

	wg.Wait()

	fmt.Println(len(frames))
	delays := make([]int, len(frames))
	for j := range delays {
		delays[j] = delay
	}

	filepath, err := runtime.SaveFileDialog(c.Ctx, runtime.SaveDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: "GIF Files (*.gif)",
				Pattern:     "*.gif;",
			},
		},
		DefaultFilename:      "animated.gif",
		Title:                "Save Animated GIF",
		CanCreateDirectories: true,
	})

	if err != nil {
		log.Fatalf("Error selecting the destination file %s", err)
	}

	file, err := os.Create(filepath)

	if err != nil {
		log.Fatalf("Error opening the destination file %s", err)
	}

	defer file.Close()

	if err := gif.EncodeAll(file, &gif.GIF{Image: frames, Delay: delays, LoopCount: 0}); err != nil {
		log.Printf("Error encoding output into animated gif :%s", err)
	}

	fmt.Println("Done")

	b, _ := os.ReadFile(filepath)

	if err != nil {
		panic(err)
	}

	str := base64.StdEncoding.EncodeToString(b)
	str = fmt.Sprintf("data:image/png;base64,%s", str)

	return filepath
	// animated := &gif.GIF{
	// 	Image: []*image.Paletted{},
	// 	Delay: []int{30, 30},
	// 	LoopCount: 0,
	// }

	// for _, layer := range layers {
	// 	file, err := os.Open(layer)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	decoded, err := png.Decode(file)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	img := image.NewPaletted(image.Rect(0, 0, 1500, 1500), palette.Plan9)
	// 	draw.Draw(img, img.Bounds(), decoded, image.ZP, draw.Src)
	// 	animated.Image = append(animated.Image, img)
	// }

	// output, err := os.OpenFile("./animated.gif", os.O_WRONLY|os.O_CREATE, 0600)

	// if err != nil {
	// 	panic(err)
	// }

	// gif.EncodeAll(output, animated)
	// fmt.Println("Running GenerateCollectionGif")
}

func (c *CollectionService) UploadCollection(dir string) {
	// var jobs []lib.Job

	// filepaths, _ := c.PrepareUploadsFromDirectory(dir)

	// for i, filepath := range filepaths {
	// 	jobs = append(jobs, lib.Job{Id: i, Config: FileItem{filepath: filepath}})
	// }

	var (
		wg        sync.WaitGroup
		// completed = 0
		storage   = lib.GetStorage()
	)

	fmt.Println(dir)

	// Create empty directory
	// dirpath := path.Join(c.DocsDirecotry, "Collection Name")
	// err := os.Mkdir(dirpath, os.ModeDir|os.ModePerm)
	// if err != nil {
	// 	panic(err)
	// }

	wg.Add(1)

	go func() {
		defer wg.Done()
		str, err := storage.PinFile(dir)
		if err != nil {
			panic(err)
		}
		fmt.Println(str)
	}()

	// go func() {
	// 	defer wg.Done()
	// 	lib.Batch(c.Ctx, 0, jobs, func(ctx context.Context, id int, job lib.Job) {
	// 		fileitem := job.Config.(FileItem)

	// 		defer func() {
	// 			completed++
	// 			data := map[string]string{"File": fileitem.filepath}
	// 			runtime.EventsEmit(ctx, "collection.item.uploaded", data)
	// 			fmt.Println(fileitem.filepath)
	// 		}()

	// 		h, _ := storage.PinFile(fileitem.filepath)
	// 		fmt.Println(h)
	// 	})
	// }()

	wg.Wait()
	fmt.Println("Done")
}

func (c *CollectionService) PrepareUploadsFromDirectory(dir string) ([]string, error) {
	// imagesdir := path.Join(dir, "images'")
	// metadatadir := path.Join(dir, "metadata")
	// if _, err := os.Stat(imagesdir); err != nil {
	// 	panic(err)
	// 	return false
	// }
	// if _, err := os.Stat(metadatadir); err != nil {
	// 	panic(err)
	// 	return false
	// }

	fmt.Println(dir)
	var uploads []string

	imageexp := regexp.MustCompile(ImageTypeExpression)
	jsonexp := regexp.MustCompile(JsonTypeExpression)

	err := filepath.Walk(
		dir,
		func(filepath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() == false {
				if imageexp.MatchString(info.Name()) || jsonexp.MatchString(info.Name()) {
					uploads = append(uploads, filepath)
				}
			}

			return nil
		})
	if err != nil {
		// lib.ErrorModal(c.Ctx, "Collection cannot be uploaded from directory", err.Error())
		return nil, err
	}

	return uploads, nil
}
