package mysqlutil

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sanity-io/litter"

	_ "github.com/go-sql-driver/mysql"

	"github.com/leizongmin/go/sqlutil"
)

func TestTable(t *testing.T) {
	sqlutil.EnableDebug()
	conn := sqlutil.MustOpenWithOptions("mysql", ConnectionOptions{
		Host:       "localhost",
		Port:       3306,
		User:       "root",
		Password:   "",
		Database:   "mysql",
		Charset:    "utf8mb4",
		Timezone:   "+8:00",
		ParseTime:  true,
		AutoCommit: true,
		Params:     nil,
	})
	defer conn.Close()

	list := make([]struct {
		Host string `db:"host"`
		DB   string `db:"db"`
		User string `db:"user"`
	}, 0)
	ok := sqlutil.QueryMany(conn, &list, Table("db").Select("host", "db", "user").Where("host='%'").Build())
	assert.True(t, ok)
	litter.Dump(list)

	item := struct {
		XX string `db:"xx"`
	}{}
	ok2 := sqlutil.QueryOne(conn, &item, Custom("SELECT ? AS xx", `"a"`))
	assert.True(t, ok2)
	litter.Dump(item)
}
