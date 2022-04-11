package lib

import (
	"encoding/json"
	"os"
)

type Metadata struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Image       string          `json:"image"`
	ExternalURL string          `json:"external_url"`
	Attributes  []MetadataAttribute `json:"attributes"`
}

type MetadataAttribute struct {
	Type  string `json:"trait_type"`
	Value string `json:"value"`
}

func GenerateMetadata(metadata Metadata, output string) bool {
	b, err := json.MarshalIndent(metadata, "", "  ")

	if err != nil {
		return false
	}

	err = os.WriteFile(output, b, os.ModePerm)

	if err != nil {
		return false
	}

	return true
}
