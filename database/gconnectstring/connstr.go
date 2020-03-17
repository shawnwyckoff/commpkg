package gconnectstring

import (
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	. "github.com/shawnwyckoff/gpkg/database/gdriver"
	"net"
	"strconv"
	"upper.io/db.v3/mongo"
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/sqlite"
)

type ConnectInfo struct {
	Host     string // host without port (e.g. localhost) or path to unix domain socket directory (e.g. /private/tmp)
	Port     uint16
	Database string
	User     string
	Password string
	Options  map[string]string
}

func Parse(dvr Driver, s string) (*ConnectInfo, error) {
	res := ConnectInfo{}
	res.Options = make(map[string]string)

	switch dvr {
	case DriverSQLite:
		slconf, err := sqlite.ParseURL(s)
		if err != nil {
			return nil, err
		}
		res.Host = ""
		res.Port = 0
		res.User = ""
		res.Password = ""
		res.Database = slconf.Database
		res.Options = slconf.Options
		return &res, nil

	case DriverMySQL:
		myconf, err := mysql.ParseURL(s)
		if err != nil {
			return nil, err
		}
		host, portstr, err := net.SplitHostPort(myconf.Host)
		if err != nil {
			return nil, err
		}
		port, err := strconv.Atoi(portstr)
		if err != nil {
			return nil, err
		}
		res.Host = host
		res.Port = uint16(port)
		res.Database = myconf.Database
		res.User = myconf.User
		res.Password = myconf.Password
		res.Options = myconf.Options
		return &res, nil

	case DriverPostgreSQL:
		pgconf, err := pgx.ParseConnectionString(s)
		res.Host = pgconf.Host
		res.Port = pgconf.Port
		res.User = pgconf.User
		res.Password = pgconf.Password
		res.Options = nil
		if err != nil {
			return nil, err
		}
		return &res, nil

	case DriverTiDB:
		return Parse(DriverMySQL, s)

	case DriverMongoDB:
		mgconf, err := mongo.ParseURL(s)
		if err != nil {
			return nil, err
		}
		host, portstr, err := net.SplitHostPort(mgconf.Host)
		if err != nil {
			return nil, err
		}
		port, err := strconv.Atoi(portstr)
		if err != nil {
			return nil, err
		}
		res.Host = host
		res.Port = uint16(port)
		res.User = mgconf.User
		res.Password = mgconf.Password
		res.Database = mgconf.Database
		res.Options = mgconf.Options
		return &res, nil
	}

	return nil, errors.Wrap(ErrUnsupportedDriver, dvr.String())
}

func (ci *ConnectInfo) Build(dvr Driver) (string, error) {
	return "", nil
	switch dvr {
	case DriverMySQL:
		//return fmt.Sprintf("%s:%s@%s/%s?%s", ci.User, ci.Password, ci.Host, ci.Database, ci.Option)

	}
	return "", nil
}
