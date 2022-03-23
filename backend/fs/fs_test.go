package fs_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/varlyapp/varlyapp/backend/fs"
)

func TestGetApplicationDocumentsDirectory(t *testing.T) {
	actual := fs.GetApplicationDocumentsDirectory()

	fmt.Println(actual)
	if strings.HasSuffix(actual, "app.varlyapp.app/Documents") != true {
		t.Fail()
	}
}
