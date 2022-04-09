package img

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"

	"github.com/disintegration/imaging"
)

func MakePreview(layers []string, width int, height int) (string, error) {
	preview := imaging.New(width, height, color.NRGBA{0, 0, 0, 0})

	// @todo this can fail if file doesn't exist, add error check
	for _, layer := range layers {
		img, _ := imaging.Open(layer)
		preview = imaging.Overlay(preview, img, image.Pt(0, 0), 1)
	}

	// imaging.Save(preview, "/Users/selvinortiz/Desktop/preview-before.png")

	preview = imaging.Resize(preview, 512, 0, imaging.Lanczos)

	// imaging.Save(preview, "/Users/selvinortiz/Desktop/preview-after.png")

	var buff bytes.Buffer

	png.Encode(&buff, preview)
	encoded := base64.StdEncoding.EncodeToString(buff.Bytes())
	encoded = fmt.Sprintf("data:image/png;base64,%s", encoded)

	return encoded, nil
}
