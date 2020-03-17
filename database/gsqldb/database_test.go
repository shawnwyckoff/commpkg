package gsqldb

import (
	"github.com/shawnwyckoff/gpkg/database/gdriver"
	"os"
	"testing"
)

func _createdatabase_test(drv gdriver.Driver) error {
	c, err := NewConn(drv, "test.db", DbNotExistOptCreateNew)
	if err != nil {
		return err
	}

	defer func() {
		_ = c.Close()
		_ = os.Remove("test.db")
	}()
	if err := c.CreateDatabaseIfNotExists("test-db"); err != nil {
		return err
	}
	return nil
}

func TestConn_CreateDatabase(t *testing.T) {
	if err := _createdatabase_test(gdriver.DriverSQLite); err != nil {
		t.Fatal(err)
	}
}
