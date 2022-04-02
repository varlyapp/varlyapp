package img

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"

	"github.com/disintegration/imaging"
)

func MakePreview(layers []string) (string, error) {
	var preview *image.NRGBA

	for _, layer := range layers {
		img, _ := imaging.Open(layer)
		preview = imaging.Overlay(preview, img, image.Pt(0, 0), 1)
	}

	preview = imaging.CropCenter(preview, 512, 512)

	var buff bytes.Buffer

	png.Encode(&buff, preview)
	encoded := base64.StdEncoding.EncodeToString(buff.Bytes())
	encoded = fmt.Sprintf("data:image/png;base64,%s", encoded)

	return encoded, nil
}
