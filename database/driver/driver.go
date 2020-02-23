package driver

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	ErrUnsupportedDriver = errors.New("Unsupported Database Driver")
)

type Driver int

// Don't support old databases such like MSSQL, Oracle, Access...
// Old databases will break compatibility.
const (
	DriverSQLite Driver = iota + 0
	DriverMySQL
	DriverPostgreSQL
	DriverTiDB
	DriverCockroachDB
	DriverMongoDB
)

func (dvr Driver) String() string {
	switch dvr {
	case DriverSQLite:
		return "SQLite"
	case DriverMySQL:
		return "MySQL"
	case DriverPostgreSQL:
		return "PostgreSQL"
	case DriverTiDB:
		return "TiDB"
	case DriverCockroachDB:
		return "CockroachDB"
	case DriverMongoDB:
		return "MongoDB"
	}
	return fmt.Sprintf("Unknown Database Driver (%d)", dvr)
}
