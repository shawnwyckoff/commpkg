package main

import (
	"fmt"
	"github.com/shawnwyckoff/gpkg/database/gdriver"
	"github.com/shawnwyckoff/gpkg/database/gsqldb"
	"os"
)

func main() {
	c, err := gsqldb.NewConn(gdriver.DriverSQLite, "test.db", gsqldb.DbNotExistOptCreateNew)
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
