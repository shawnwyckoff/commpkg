package excel

import (
	"github.com/shawnwyckoff/gpkg/sys/fs"
	"testing"
)

func TestMemDoc_ToXlsx(t *testing.T) {
	s := "test.xlsx"
	xd, err := OpenPath(s)
	if err != nil {
		t.Error(err)
	}

	b, err := xd.ToMemDoc().ToXlsx()
	if err != nil {
		t.Error(err)
	}
	fs.BytesToFile(b, s + "2.xlsx")
}