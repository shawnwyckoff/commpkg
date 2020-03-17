package gsqldb

import (
	"github.com/shawnwyckoff/gpkg/apputil/gtest"
	"github.com/shawnwyckoff/gpkg/database/gdriver"
	"testing"
)

type sheets struct {
	K string `xorm:"varchar(256) pk not null 'k'"`
	V string `xorm:"JSON 'v'"`
}

func dialTestDB() (*SqlDB, error) {
	return Dial(gdriver.MySQL, "root:msq%!888@tcp(192.168.9.20:3306)/whale?charset=utf8")
}

func TestSqlDB_Tables(t *testing.T) {
	s, err := dialTestDB()
	gtest.Assert(t, err)
	defer s.Close()

	tables, err := s.Tables()
	gtest.Assert(t, err)

	t.Log(tables)
}

func TestSqlDB_SelectOne(t *testing.T) {
	s, err := dialTestDB()
	gtest.Assert(t, err)
	defer s.Close()

	out := sheets{K:"本季度2"}
	ok, err := s.SelectOne(&out)
	gtest.Assert(t, err)
	t.Log(ok)
	t.Log(out)
}

func TestSqlDB_SelectAll(t *testing.T) {
	s, err := dialTestDB()
	gtest.Assert(t, err)
	defer s.Close()

	out := make([]sheets, 0)
	err = s.SelectAll(&out)
	gtest.Assert(t, err)
	t.Log(out)
}

func TestSqlDB_UpsertOne(t *testing.T) {
	s, err := dialTestDB()
	gtest.Assert(t, err)
	defer s.Close()

	newRecord := sheets{
		K: "本季度2",
		V: `{"name":"tom", "age":22}`,
	}
	n, err := s.UpsertOne(newRecord, &sheets{K:newRecord.K})
	gtest.Assert(t, err)
	t.Log(n)
}

func TestSqlDB_Exist(t *testing.T) {
	s, err := dialTestDB()
	gtest.Assert(t, err)
	defer s.Close()

	exist, err := s.Exist(&sheets{K:"本季度"})
	gtest.Assert(t, err)
	t.Log(exist)
}

func TestSqlDB_Remove(t *testing.T) {
	s, err := dialTestDB()
	gtest.Assert(t, err)
	defer s.Close()

	n, err := s.Remove(&sheets{K:"本季度2"})
	gtest.Assert(t, err)
	t.Log(n)
}