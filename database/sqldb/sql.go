package sqldb

import (
	"fmt"
	"github.com/pkg/errors"
	. "github.com/shawnwyckoff/gpkg/database/driver"
)

// https://github.com/samonzeweb/godb

// SQL语句末尾要加分号
// 单引号双引号

type SQLBuilder struct {
}

func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{}
}

func (b *SQLBuilder) CreateDatabase(dvr Driver, database string) (string, error) {
	return fmt.Sprintf("CREATE DATABASE \"%s\";", database), nil
}

// In PostgreSQL，CREATE DATABASE IF NOT EXISTS is NOT supported
// Reference: https://stackoverflow.com/questions/44511958/python-postgresql-create-database-if-not-exists-is-error/44512503#44512503
func (b *SQLBuilder) CreateDatabaseIfNotExists(dvr Driver, database string) (string, error) {
	if dvr == DriverPostgreSQL {
		return "", errors.Errorf("%s doesn't support CREATE DATABASE IF NOT EXISTS statement", dvr.String())
	}
	if dvr == DriverSQLite || dvr == DriverMySQL ||
		dvr == DriverTiDB || dvr == DriverCockroachDB {
		return fmt.Sprintf("CREATE DATABASE IF NOT EXISTS \"%s\";", database), nil
	}
	return "", errors.Errorf("CreateDatabaseIfNotExists does NOT suppoted driver %s", dvr.String())
}

func (b *SQLBuilder) DropDatabase(dvr Driver, database string) (string, error) {
	if dvr == DriverSQLite || dvr == DriverMySQL ||
		dvr == DriverTiDB || dvr == DriverPostgreSQL ||
		dvr == DriverCockroachDB {
		return fmt.Sprintf("DROP DATABASE \"%s\";", database), nil
	}

	return "", errors.Errorf("DropDatabase does NOT suppoted driver %s", dvr.String())
}

func (b *SQLBuilder) DropDatabaseIfExists(dvr Driver, database string) (string, error) {
	if dvr == DriverSQLite {
		return "", errors.Errorf("For SQLite every .db file is an unique database, can't drop database with SQL")
	}
	if dvr == DriverMySQL || dvr == DriverTiDB ||
		dvr == DriverPostgreSQL || dvr == DriverCockroachDB {
		// In PostgreSQL，DROP DATABASE IF EXISTS 'dbname'; is wrong
		// DROP DATABASE IF EXISTS "dbname"; is right
		return fmt.Sprintf("DROP DATABASE IF EXISTS \"%s\";", database), nil
	}
	return "", errors.Errorf("DropDatabaseIfExists does NOT suppoted driver %s", dvr.String())
}

// In PostgreSQL，DROP TABLE IF EXISTS is supported
func (b *SQLBuilder) DropTableIfExists(dvr Driver, table string) (string, error) {
	if dvr == DriverSQLite || dvr == DriverMySQL ||
		dvr == DriverTiDB || dvr == DriverPostgreSQL ||
		dvr == DriverCockroachDB {
		return fmt.Sprintf("DROP TABLE IF EXISTS \"%s\";", table), nil
	}
	return "", errors.Errorf("DropTableIfExists does NOT suppoted driver %s", dvr.String())
}

func (b *SQLBuilder) CountTable(dvr Driver, table string) (string, error) {
	if dvr == DriverSQLite || dvr == DriverMySQL ||
		dvr == DriverTiDB || dvr == DriverPostgreSQL ||
		dvr == DriverCockroachDB {
		return fmt.Sprintf("SELECT COUNT(*) FROM \"%s\";", table), nil
	}
	return "", errors.Errorf("CountTable does NOT suppoted driver %s", dvr.String())
}

// If exists return string 'yes', otherwise return 'no'
// For SQLite, one disk file is the unique database, so check the disk file
func (b *SQLBuilder) CheckDatabaseExists(dvr Driver, database string) (string, error) {
	if dvr == DriverPostgreSQL || dvr == DriverCockroachDB {
		return fmt.Sprintf("SELECT CASE WHEN(EXISTS (SELECT * FROM pg_database WHERE datname='%s')) THEN 'yes' ELSE 'no' END;", database), nil
	}
	if dvr == DriverMySQL || dvr == DriverTiDB {
		return fmt.Sprintf("SELECT IF(EXISTS (SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = '%s'), 'yes','no');", database), nil
	}
	if dvr == DriverSQLite {
		return "", errors.Errorf("For SQLite every .db file is an unique database, can't check it with SQL")
	}

	return "", errors.Errorf("CheckDatabaseExists does NOT suppoted driver %s", dvr.String())

}

func (b *SQLBuilder) ShowDatabases(dvr Driver) (string, error) {
	if dvr == DriverPostgreSQL || dvr == DriverCockroachDB {
		return "SELECT datname FROM pg_database;", nil
	}
	if dvr == DriverMySQL || dvr == DriverTiDB {
		return "SHOW DATABASES;", nil
	}
	return "", errors.Errorf("ShowDatabases does NOT suppoted driver %s", dvr.String())
}

func (b *SQLBuilder) ShowTables(dvr Driver) (string, error) {
	if dvr == DriverPostgreSQL || dvr == DriverCockroachDB {
		return "SELECT * FROM pg_catalog.pg_tables;", nil
	}
	if dvr == DriverMySQL || dvr == DriverTiDB {
		return "SHOW TABLES;", nil
	}
	if dvr == DriverSQLite {
		return ".tables", nil
	}
	return "", errors.Errorf("ShowTables does NOT suppoted driver %s", dvr.String())
}
