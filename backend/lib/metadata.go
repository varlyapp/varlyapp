package lib

import (
	"encoding/json"
	"os"
)

type Metadata struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Image       string          `json:"image"`
	Traits      []MetadataTrait `json:"attributes"`

	// Solana JSON schema
	Collection   MetadataCollection `json:"collection"`
	Symbol       string             `json:"symbol"`
	AnimationURL string             `json:"animatin_url"`
	ExternalURL  string             `json:"external_url"`

	SellerFeeBasisPoints string `json:"seller_fee_basis_points"`
	// Other properties
	// properties[files[], category, creators[]]
}

type MetadataTrait struct {
	Type  string `json:"trait_type"`
	Value string `json:"value"`
}

type MetadataCollection struct {
	Name   string `json:"name"`
	Family string `json:"family"`
}

func GenerateMetadata(metadata Metadata, filepath string) error {
	b, err := json.MarshalIndent(metadata, "", "  ")

	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, b, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
