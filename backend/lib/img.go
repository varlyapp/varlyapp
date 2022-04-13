package lib

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"

	"github.com/disintegration/imaging"
)

func MakePreview(layers []string, width int, height int, size int) (string, error) {
	preview := imaging.New(width, height, color.NRGBA{0, 0, 0, 0})

	// @todo this can fail if file doesn't exist, add error check
	for _, layer := range layers {
		img, err := imaging.Open(layer)
		if err != nil {
			return "", err
		}
		preview = imaging.Overlay(preview, img, image.Pt(0, 0), 1)
	}

	imaging.Save(preview, "/Users/selvinortiz/Desktop/preview-before.png")

	preview = imaging.Resize(preview, size, 0, imaging.Lanczos)

	imaging.Save(preview, "/Users/selvinortiz/Desktop/preview-after.png")

	var buff bytes.Buffer

	png.Encode(&buff, preview)
	encoded := base64.StdEncoding.EncodeToString(buff.Bytes())
	encoded = fmt.Sprintf("data:image/png;base64,%s", encoded)

	return encoded, nil
}

func GeneratePNG(layers []string, filepath string, width int, height int) error {
	png := imaging.New(width, height, color.NRGBA{0, 0, 0, 0})

	for _, layer := range layers {
		img, err := imaging.Open(layer)
		if err != nil {
			return err
		}
		png = imaging.Overlay(png, img, image.Pt(0, 0), 1)
	}
	err := imaging.Save(png, filepath)
	if err != nil {
		return err
	}

	return nil
}
