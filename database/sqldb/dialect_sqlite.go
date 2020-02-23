package sqldb

import (
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/commpkg/sys/fs"
	"os"
)

func _sqliteIsDatabaseEixsts(name string) (bool, error) {
	pi, err := fs.GetPathInfo(name)
	if err != nil {
		return false, err
	}
	if pi.IsFolder {
		return false, nil
	}
	return pi.Exist, nil
}

func _sqliteRemoveDatabase(name string) error {
	if !fs.FileExits(name) {
		return errors.Errorf("SQLite database file '%s' not exists", name)
	}
	return os.Remove(name)
}
