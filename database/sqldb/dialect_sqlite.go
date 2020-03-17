package sqldb

import (
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/gpkg/sys/gfs"
	"os"
)

func _sqliteIsDatabaseEixsts(name string) (bool, error) {
	pi, err := gfs.GetPathInfo(name)
	if err != nil {
		return false, err
	}
	if pi.IsFolder {
		return false, nil
	}
	return pi.Exist, nil
}

func _sqliteRemoveDatabase(name string) error {
	if !gfs.FileExits(name) {
		return errors.Errorf("SQLite database file '%s' not exists", name)
	}
	return os.Remove(name)
}
