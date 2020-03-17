package main

import (
	"fmt"
	"github.com/shawnwyckoff/gpkg/database/gdriver"
	"github.com/shawnwyckoff/gpkg/database/gsqldb_old_to_remove"
	"os"
)

func main() {
	c, err := gsqldb_old_to_remove.NewConn(gdriver.SQLite, "test.db", gsqldb_old_to_remove.DbNotExistOptCreateNew)
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
