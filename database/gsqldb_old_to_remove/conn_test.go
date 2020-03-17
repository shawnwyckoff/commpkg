package gsqldb_old_to_remove

import (
	. "github.com/shawnwyckoff/gpkg/database/gdriver"
	"os"
	"testing"
)

func TestNewConn(t *testing.T) {
	c, err := NewConn(SQLite, "test.db", DbNotExistOptCreateNew)
	if err != nil {
		t.Fatal(err)
	}
	c.Close()
	os.Remove("test.db")
}
