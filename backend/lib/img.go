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
	// started := time.Now()
	preview := imaging.New(size, size, color.NRGBA{0, 0, 0, 0})
	// fmt.Printf("Created placeholder preview: %s\n", time.Since(started))
	for _, layer := range layers {
		img, err := imaging.Open(layer)
		if err != nil {
			return "", err
		}
		img = imaging.Resize(img, size, 0, imaging.Box)
		preview = imaging.Overlay(preview, img, image.Pt(0, 0), 1)
		// fmt.Printf("Created overlay %d: %s\n", index, time.Since(started))
	}

	// imaging.Save(preview, "/Users/selvinortiz/Desktop/preview-before.png")

	// preview = imaging.Resize(preview, size, 0, imaging.Lanczos)
	// fmt.Printf("Resized preview: %s\n", time.Since(started))

	// imaging.Save(preview, "/Users/selvinortiz/Desktop/preview-after.png")

	var buff bytes.Buffer

	png.Encode(&buff, preview)
	encoded := base64.StdEncoding.EncodeToString(buff.Bytes())
	encoded = fmt.Sprintf("data:image/png;base64,%s", encoded)

	// fmt.Printf("Encoded preview: %s\n", time.Since(started))

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

func GenerateFrame(layers []string, filepath string, width int, height int, size int) error {
	png := imaging.New(width, height, color.NRGBA{0, 0, 0, 0})

	for _, layer := range layers {
		img, err := imaging.Open(layer)
		if err != nil {
			return err
		}
		img = imaging.Resize(img, size, 0, imaging.Box)
		png = imaging.Overlay(png, img, image.Pt(0, 0), 1)
	}

	err := imaging.Save(png, filepath)
	if err != nil {
		return err
	}

	return nil
}
