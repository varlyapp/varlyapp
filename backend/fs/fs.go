package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

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

func GetApplicationDocumentsDirectory(paths ...string) (string, error) {
	dir, _ := os.UserConfigDir()

	path := fmt.Sprintf("%s/varlyapp/Documents/%s", dir, strings.Join(paths, "/"))

	err := os.MkdirAll(path, os.ModePerm);

	if err != nil {
		return "", err
	}

	return path, nil
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

func UploadCollection(key string, secret string) string {
	storage := pinata.Pinata{Apikey: key, Secret: secret}
	cid, err := storage.PinFile("/Users/selvinortiz/Downloads/varly-collection.png")
	if err != nil {
		return fmt.Sprintln(err)
	}
	return cid
}

func Include(path string) string {

	content, _ := os.ReadFile(fmt.Sprint("/assets/icons", path))

	fmt.Println(string(content))
	return string(content)
}
