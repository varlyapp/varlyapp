package types

type VariantFileType string

const (
	VariantFileTypePng  VariantFileType = "PNG"
	VariantFileTypeTiff                 = "TIFF"
	VariantFileTypeGif                  = "GIF"
)

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
