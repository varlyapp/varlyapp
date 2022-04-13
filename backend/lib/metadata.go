package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Metadata struct {
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	AnimationURL string `json:"animatin_url"`
	ExternalURL  string `json:"external_url"`

	DNA      string `json:"dna"`
	Edition  int    `json:"edition"`
	Artist   string `json:"artist"`
	Compiler string `json:"compiler"`
	Date     int64  `json:"date"`

	// Solana JSON schema
	Collection MetadataCollection `json:"collection"`

	SellerFeeBasisPoints string `json:"seller_fee_basis_points"`
	// Other properties
	// properties[files[], category, creators[]]

	// List of traits
	Traits []MetadataTrait `json:"attributes"`
}

type MetadataTrait struct {
	Type  string `json:"trait_type"`
	Value string `json:"value"`
}

type MetadataCollection struct {
	Name   string `json:"name"`
	Family string `json:"family"`
}

type MetadataConfig struct {
	CollectionName        string
	CollectionDescription string
	CollectionSymbol      string
	CollectionBaseURI     string
	Artist                string
	DNA                   string
	Name                  string
	Edition               int
	Layers                []string
	Image                 string
}

func LayersToMetadataTraits(layers []string) []MetadataTrait {
	var traits []MetadataTrait

	for _, layer := range layers {
		dir, file := filepath.Split(layer)
		traits = append(traits, MetadataTrait{
			Type:  filepath.Base(dir),
			Value: strings.Replace(file, ".png", "", 1),
		})
	}

	return traits
}

func ParsePlaceholder(s string, placeholder string, replacement string) string {
	return strings.ReplaceAll(s, placeholder, replacement)
}

func ParseImageURI(path string, baseURI string) (string, string) {
	_, file := filepath.Split(path)

	return fmt.Sprintf("%s%s", baseURI, file), strings.Replace(path, ".png", ".json", 1)
}

func GenerateMetadata(config MetadataConfig) error {
	name := ParsePlaceholder(config.CollectionName, "{{edition}}", fmt.Sprintf("%d", config.Edition))
	traits := LayersToMetadataTraits(config.Layers)
	imageURI, jsonFilepath := ParseImageURI(config.Image, config.CollectionBaseURI)

	metadata := Metadata{
		Name:        name,
		Description: "",
		Image:       imageURI,
		Artist:      config.Artist,
		Edition:     config.Edition,
		DNA:         config.DNA,
		Date:        time.Now().Unix(),

		// Solana JSON schema
		Collection:   MetadataCollection{},
		Symbol:       "{{SYMBOL}}",
		AnimationURL: "",
		ExternalURL:  "",

		SellerFeeBasisPoints: "",
		Compiler:             "Varly.app",

		// List of Traits
		Traits: traits,
	}
	b, err := json.MarshalIndent(metadata, "", "  ")

	if err != nil {
		return err
	}

	err = os.WriteFile(jsonFilepath, b, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
