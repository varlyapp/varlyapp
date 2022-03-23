package nft_test

import (
	"testing"

	"github.com/varlyapp/varlyapp/backend/nft"
)

func TestGeneratePng(t *testing.T) {
	actual := nft.GeneratePNG([]string{
		"/Users/selvinortiz/Desktop/Varly Example/Background/Yellow.png",
		"/Users/selvinortiz/Desktop/Varly Example/Icon/BTT.png",
		"/Users/selvinortiz/Desktop/Varly Example/Mark/Octagon.png",
	}, "/Users/selvinortiz/Desktop/Build/Example.png", 3000, 3000,
	)

	if actual != true {
		t.Fail()
	}
}
