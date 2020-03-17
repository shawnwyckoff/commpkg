package hdd

import (
	"encoding/json"
	"github.com/shawnwyckoff/gpkg/xsystils/xsys/xproc"
	"log"
	"testing"
)

func TestGetVolumeInfo(t *testing.T) {
	_, _, mydir, err := xproc.SelfPath()
	if err != nil {
		t.Error(err)
		return
	}

	vi, err := GetVolumeInfo(mydir)
	if err != nil {
		t.Error(err)
		return
	}
	buf, err := json.Marshal(vi)
	if err != nil {
		t.Error(err)
		return
	}
	log.Print(string(buf))
}
