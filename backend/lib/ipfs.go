package lib

import (
	"fmt"

	"github.com/wabarc/ipfs-pinner/pkg/pinata"
)

func UploadCollection(key string, secret string) string {
	storage := pinata.Pinata{Apikey: key, Secret: secret}
	cid, err := storage.PinFile("/Users/selvinortiz/Downloads/varly-collection.png")
	if err != nil {
		return fmt.Sprintln(err)
	}
	return cid
}
