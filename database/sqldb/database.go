package sqldb

import (
	"github.com/pkg/errors"
	. "github.com/shawnwyckoff/commpkg/database/driver"
	"github.com/shawnwyckoff/commpkg/database/sql"
	"strings"
)

type (
	DB struct {
		conn *Conn
		name string
	}

	Table struct {
		db   *DB
		name string
	}
)

func (c *Conn) ShowDatabases() ([]string, error) {
	if c.driver == DriverSQLite {
		// Sqlite3 doesn't support CreateDatabaseIfNotExists too
		return nil, errors.Wrap(ErrUnsupportedDriver, c.driver.String()+":ShowDatabases()")
	}
	sql, err := sql.NewBuilder().ShowDatabases(c.driver)
	if err != nil {
		return nil, err
	}
	ds, err := c.Query(sql)
	if err != nil {
		return nil, err
	}
	return ds.ReadStringArray()
}

func (c *Conn) DatabasesCount() (int, error) {
	dbs, err := c.ShowDatabases()
	if err != nil {
		return 0, err
	}
	return len(dbs), nil
}

// Get database name of current connection
func (c *Conn) DatabaseName() (string, error) {
	return c.connInfo.Database, nil
	// in xorm, you can get dbname using c.xormConn.Dialect().URI().DbName
	// but in jinzhu's gorm, you can't get dbname using it's api
}

// Whether database exist on connected host
func (c *Conn) DatabaseExists(name string) (bool, error) {
	// For SQLite, one disk file is the unique database, so check the disk file
	if c.driver == DriverSQLite {
		return _sqliteIsDatabaseEixsts(name)
	}

	// Check by SQL
	sql, err := sql.NewBuilder().CheckDatabaseExists(c.driver, name)
	if err != nil {
		return false, err
	}
	ds, err := c.Query(sql)
	if err != nil {
		return false, err
	}
	str, err := ds.ReadFirstString()
	if err != nil {
		return false, err
	}
	switch strings.ToLower(str) {
	case "yes":
		return true, nil
	case "no":
		return false, nil
	}
	return false, errors.Errorf("DatabaseExists: Unknown return \"%s\"", str)
}

func (c *Conn) CreateDatabaseIfNotExists(name string) error {
	if c.driver == DriverPostgreSQL {
		exist, err := c.DatabaseExists(name)
		if err != nil {
			return err
		}
		if exist {
			return nil
		}
		sql, err := sql.NewBuilder().CreateDatabase(c.driver, name)
		if err != nil {
			return err
		}
		return c.Exec(sql)
	} else if c.driver == DriverMySQL || c.driver == DriverTiDB {
		sql, err := sql.NewBuilder().CreateDatabaseIfNotExists(c.driver, name)
		if err != nil {
			return err
		}
		return c.Exec(sql)
	}

	// Sqlite3 doesn't support CreateDatabaseIfNotExists too
	return errors.Wrap(ErrUnsupportedDriver, c.driver.String()+":CreateDatabaseIfNotExists()")
}

func (c *Conn) DropDatabaseIfExists(name string) error {
	// For SQLite, one disk file is the unique database, so remove the disk file
	if c.driver == DriverSQLite {
		return _sqliteRemoveDatabase(name)
	}

	sql, err := sql.NewBuilder().DropDatabaseIfExists(c.driver, name)
	if err != nil {
		return err
	}
	return c.Exec(sql)
}

func (c *Conn) ShowTables() ([]string, error) {
	if len(c.connInfo.Database) == 0 {
		return nil, ErrCurrDatabaseIsNull
	}

	sql, err := sql.NewBuilder().ShowTables(c.driver)
	if err != nil {
		return nil, err
	}
	ds, err := c.Query(sql)
	if err != nil {
		return nil, err
	}
	return ds.ReadStringArray()
}

func (c *Conn) TablesCount() (int, error) {
	if len(c.connInfo.Database) == 0 {
		return 0, ErrCurrDatabaseIsNull
	}

	tbs, err := c.ShowTables()
	if err != nil {
		return 0, err
	}
	return len(tbs), nil
}

func (c *Conn) DB(dbName string) *DB {
	return &DB{conn: c, name: dbName}
}

func (db *DB) Table(tableName string) *Table {
	return &Table{db: db, name: tableName}
}
