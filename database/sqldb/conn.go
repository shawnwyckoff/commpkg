package sqldb

// Support databases:
// MySQL, PostgreSQL, Sqlite, TiDB, CockroachDB

// 按库的质量，排序如下
// gorm
// MySQL, PostgreSQL, Sqlite
//
// xorm
// TiDB, MSSQL, Oracle
//
// upper/db
// MongoDB

// https://github.com/jinzhu/gorm // 8000 star mysql,postgres,sqlite
// https://github.com/go-xorm/xorm // 3000 star mysql,postgres,tidb,sqlite3,mssql,oracle
// https://github.com/go-gorp/gorp // 2700 star mysql,postgres,sqlite,sqlserver,oracle
// https://github.com/upper/db // 1100 star 产品级,支持数据库为 PostgreSQL, MySQL, SQLite, MSSQL, QL and MongoDB
// https://github.com/go-pg/pg // 专注于pgsql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/commpkg/database/connstr"
	. "github.com/shawnwyckoff/commpkg/database/driver"
)

type DbNotExistOpt int

const (
	DbNotExistOptReturnError DbNotExistOpt = iota + 0
	DbNotExistOptCreateNew
)

// If connectString doesn't include database, will NOT return error
// If opt == DbNotExistOptCreateNew and database not exist, it will create one
func NewConn(dvr Driver, connectString string, opt DbNotExistOpt) (*Conn, error) {
	// Parse connect string
	ci, err := connectString.Parse(dvr, connectString)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Invalid connect string '%s'", connectString))
	}

	/*if dvr == DriverSQLite {
		if len(ci.Database) == 0 {
			return nil, errors.New("In sqlite3 database must appear in connect string")
		}
		if Dbnot
	}*/

	// Need to check whether database exists
	if len(ci.Database) > 0 && opt == DbNotExistOptCreateNew {
		tmpci := ci
		tmpci.Database = ""
		tmpconnstr, err := tmpci.Build(dvr)
		if err != nil {
			return nil, err
		}
		tmpconn, err := NewConn(dvr, tmpconnstr, DbNotExistOptReturnError)
		if err != nil {
			return nil, errors.Wrap(err, "Error in create temp manage connection")
		}
		defer tmpconn.Close()
		if err := tmpconn.CreateDatabaseIfNotExists(ci.Database); err != nil {
			return nil, errors.Wrap(err, "Error in create new")
		}
	}

	// Connect
	c := Conn{driver: dvr, connInfo: *ci}
	switch dvr {
	case DriverSQLite:
		c.gormConn, err = gorm.Open("sqlite3", connectString)
	case DriverMySQL:
		c.gormConn, err = gorm.Open("mysql", connectString)
	case DriverPostgreSQL:
		c.gormConn, err = gorm.Open("postgresql", connectString)
	case DriverTiDB:
		c.gormConn, err = gorm.Open("mysql", connectString)
	default:
		err = errors.Wrap(ErrUnsupportedDriver, dvr.String())
	}

	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Conn) Close() error {
	if c.gormConn != nil {
		return c.gormConn.Close()
	}
	return nil
}
