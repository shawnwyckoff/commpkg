package sqldb

import (
	"github.com/shawnwyckoff/commpkg/database/driver"
	"os"
	"testing"
)

func _createdatabase_test(drv driver.Driver) error {
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
	if err := _createdatabase_test(driver.DriverSQLite); err != nil {
		t.Fatal(err)
	}
}
