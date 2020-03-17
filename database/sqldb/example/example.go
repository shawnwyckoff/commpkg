package main

import (
	"fmt"
	"github.com/shawnwyckoff/gpkg/database/driver"
	"github.com/shawnwyckoff/gpkg/database/sqldb"
	"os"
)

func main() {
	c, err := sqldb.NewConn(driver.DriverSQLite, "test.db", sqldb.DbNotExistOptCreateNew)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		c.Close()
		os.Remove("test.db")
	}()
	if err := c.CreateDatabaseIfNotExists("test-db"); err != nil {
		fmt.Println(err)
		return
	}
}
