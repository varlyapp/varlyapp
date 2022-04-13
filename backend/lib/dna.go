package lib

import (
	"crypto/sha1"
	"fmt"
)

func GenerateDNA(layers []string) string {
	input := ""
	for _, layer := range layers {
		input += layer
	}
	return fmt.Sprintf("%x", sha1.Sum([]byte(input)))
}
