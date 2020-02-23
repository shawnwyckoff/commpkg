package sqldb

import (
	"github.com/shawnwyckoff/commpkg/dsa/interfaces"
)

func (c *Conn) TableExists(name string) (bool, error) {
	return c.gormConn.HasTable(name), nil
}

/*
 - model example:
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
*/
func (c *Conn) CreateTableIfNotExists(model interface{}) error {
	if len(c.connInfo.Database) == 0 {
		return ErrCurrDatabaseIsNull
	}
	name, _ := interfaces.Parse(model)
	exists, err := c.TableExists(name)
	if err != nil {
		return err
	}
	if exists {
		return err
	}
	return c.gormConn.CreateTable(model).Error
}

func (c *Conn) DropTableIfExists(name string) error {
	if len(c.connInfo.Database) == 0 {
		return ErrCurrDatabaseIsNull
	}

	sql, err := NewSQLBuilder().DropTableIfExists(c.driver, name)
	if err != nil {
		return err
	}

	return c.Exec(sql)
}

func (c *Conn) CountTable(table string) (int, error) {
	if len(c.connInfo.Database) == 0 {
		return 0, ErrCurrDatabaseIsNull
	}

	sql, err := NewSQLBuilder().CountTable(c.driver, table)
	if err != nil {
		return 0, err
	}
	ds, err := c.Query(sql)
	if err != nil {
		return 0, err
	}
	return ds.ReadInt()
}

// TODO
/*func (c *Conn) AddColumn(table string, def ColumnDefine) error {
	if len(c.connInfo.Database) == 0 {
		return ErrCurrDatabaseIsNull
	}

	return nil
}

func (c *Conn) AlterColumn(table, oldColumn string, newdef ColumnDefine) error {
	if len(c.connInfo.Database) == 0 {
		return ErrCurrDatabaseIsNull
	}

	return nil
}

func (c *Conn) DropColumn(table, name string) error {
	if len(c.connInfo.Database) == 0 {
		return ErrCurrDatabaseIsNull
	}

	return nil
}*/
