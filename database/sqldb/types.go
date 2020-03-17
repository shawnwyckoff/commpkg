package sqldb

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/gpkg/database/connect_string"
	. "github.com/shawnwyckoff/gpkg/database/driver"
)

type Conn struct {
	connInfo connect_string.ConnectInfo
	driver   Driver
	gormConn *gorm.DB
}

//type HostConn Conn

var (
	ErrDatabaseNotExist   = errors.New("Database not exist")
	ErrCurrDatabaseIsNull = errors.New("Host connected without database, current API must connect database")
)

type DataType int

const (
	DataTypeBOOL DataType = iota + 0
	DataTypeInt32
	DataTypeInt64
	DataTypeFloat32
	DataTypeFloat64
	DataTypeBigInt
	DataTypeDecimal
	DataTypeString
	DataTypeDatetime
	DataTypeBLOB
	DataTypeJSON
	DataTypeJSONB
)

/*
type ColumnDefine struct {
	Name         string
	TypeSide         DataType
	IsPrimaryKey bool
	IsIndex      bool
	NotNull      bool
}*/

type Dataset struct {
}
