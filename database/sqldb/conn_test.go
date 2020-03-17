package sqldb

import (
	. "github.com/shawnwyckoff/gpkg/database/driver"
	"os"
	"testing"
)

func TestNewConn(t *testing.T) {
	c, err := NewConn(DriverSQLite, "test.db", DbNotExistOptCreateNew)
	if err != nil {
		t.Fatal(err)
	}
	c.Close()
	os.Remove("test.db")
}
