package fs_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/varlyapp/varlyapp/backend/fs"
)

func TestGetApplicationDocumentsDirectory(t *testing.T) {
	actual, _ := fs.GetApplicationDocumentsDirectory()

	fmt.Println(actual)
	if strings.HasSuffix(actual, "varlyapp/Documents") != true {
		t.Fail()
	}
}
