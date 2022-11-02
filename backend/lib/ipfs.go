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

func GetStorage() *pinata.Pinata {
	return &pinata.Pinata{Apikey: "21e4c90e97570a749d30", Secret: "ef7607d9ea3e66fbe1face8740a89fb0422b25e4fadefaf33039b2c8b02c67b2"}
}

/*
API Key: 21e4c90e97570a749d30
 API Secret: ef7607d9ea3e66fbe1face8740a89fb0422b25e4fadefaf33039b2c8b02c67b2
 JWT: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mb3JtYXRpb24iOnsiaWQiOiIyZjM2ZWJiNy00NTg0LTRiZmUtODI1OS1lNDM4ZmM4MTRiY2MiLCJlbWFpbCI6InNlbHZpbkBzZWx2aW4uZGV2IiwiZW1haWxfdmVyaWZpZWQiOnRydWUsInBpbl9wb2xpY3kiOnsicmVnaW9ucyI6W3siaWQiOiJOWUMxIiwiZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjF9XSwidmVyc2lvbiI6MX0sIm1mYV9lbmFibGVkIjpmYWxzZX0sImF1dGhlbnRpY2F0aW9uVHlwZSI6InNjb3BlZEtleSIsInNjb3BlZEtleUtleSI6IjIxZTRjOTBlOTc1NzBhNzQ5ZDMwIiwic2NvcGVkS2V5U2VjcmV0IjoiZWY3NjA3ZDllYTNlNjZmYmUxZmFjZTg3NDBhODlmYjA0MjJiMjVlNGZhZGVmYWYzMzAzOWIyYzhiMDJjNjdiMiIsImlhdCI6MTY1MDg0NDYxOH0.VvhQTcWXsGaxD8xmmhh1xIA4kR9KLwcfoDTP2qALLWo
*/
