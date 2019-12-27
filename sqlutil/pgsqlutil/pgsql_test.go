package pgsqlutil

import (
	"testing"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"

	"github.com/leizongmin/go/sqlutil"
)

func TestTable(t *testing.T) {
	sqlutil.EnableDebug()
	conn := sqlutil.MustOpenWithOptions("postgres", ConnectionOptions{
		Host:     "localhost",
		Port:     5432,
		User:     "glen",
		Password: "",
		Database: "postgres",
		SSLMode:  "disable",
	})
	defer conn.Close()

	list := make([]struct {
		DatName string `db:"datname"`
		DatDba  string `db:"datdba"`
	}, 0)
	ok := sqlutil.FindMany(conn, &list, Table("pg_database").Select("datname", "datdba").Where("encoding>?", 0).Build())
	assert.True(t, ok)
	litter.Dump(list)

	item := struct {
		XX string `db:"xx"`
		YY int    `db:"yy"`
	}{}
	ok2 := sqlutil.FindOne(conn, &item, Custom("SELECT ? AS xx, ? AS yy", `"a"`, 12345))
	assert.True(t, ok2)
	litter.Dump(item)
}
