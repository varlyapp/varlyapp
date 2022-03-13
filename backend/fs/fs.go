package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/wabarc/ipfs-pinner/pkg/pinata"
)

const FileTypeExpression = `.(png|jpg|jpeg)$`
const RarityExpression = `(.*)#([.0-9]+).(png|jpg|jpeg)$`
const DefaultRarity int64 = 100

type Layers struct {
	Images map[string][]string
}

type Layer struct {
	Name   string
	Item   string
	Weight int64
}

type CollectionConfig struct {
	Layers map[string][]Layer
}

// ReadLayers reads a directory with images represending collection traits
// Each directory is considered a trait
// Each image within a directory is considered a variant of that trait
func ReadLayers(dir string) (CollectionConfig, error) {
	return ReadDirectory(dir)
}

// ReadsDirectory reads a direcotory into a CollectionConfig
// Items within the directory are split into Layers
func ReadDirectory(dir string) (CollectionConfig, error) {
	config := CollectionConfig{Layers: make(map[string][]Layer)}
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
					var weight int64

					if rarityX.MatchString(path) {
						weight, _ = strconv.ParseInt(rarityX.ReplaceAllString(path, "$2"), 10, 64)
					}

					if weight <= 0 {
						weight = DefaultRarity
					}

					config.Layers[lastDirectory] = append(config.Layers[lastDirectory], Layer{Name: info.Name(), Item: path, Weight: int64(weight)})
				}
			}

			return nil
		})

	return config, err
}

func UploadCollection() string {
	storage := pinata.Pinata{Apikey: "421dfc8596c951b2e35b", Secret: "f780d6362e442ee1441f8c786581c223f81a4d4af6ea5351821ab26b59f520bf"}
	cid, err := storage.PinFile("/Users/selvinortiz/Downloads/varly-collection.png")
	if err != nil {
		return fmt.Sprintln(err)
	}
	return cid
}
