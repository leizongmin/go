package sqlConnection

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestBuildDataSourceString(t *testing.T) {
	str := BuildDataSourceString(Options{
		Host:       "127.0.0.1",
		Port:       3306,
		User:       "root",
		Password:   "123",
		Database:   "xxx",
		Charset:    "utf8mb4",
		Timezone:   "+8:00",
		ParseTime:  true,
		AutoCommit: true,
		Params:     nil,
	})
	fmt.Println(str)
	assert.Equal(t, "root:123@tcp(127.0.0.1:3306)/xxx?&parseTime=true&loc=Local&charset=utf8mb4&time_zone=%27%2B8%3A00%27&autocommit=true", str)
}

func TestFindMany(t *testing.T) {
	EnableDebug()
	db, err := sqlx.Open("mysql", BuildDataSourceString(Options{Database: "mysql"}))
	assert.NoError(t, err)
	fmt.Printf("%+v\n", db)

	count, ok := FindCount(db, "SELECT COUNT(*) AS count FROM user")
	fmt.Println(count, ok)

	var list []Row
	var tx Tx
	tx = db.MustBegin()

	ok = FindMany(tx, &list, "SHOW TABLES")
	fmt.Printf("%v %+v\n", ok, list)

	rows, err := db.Queryx("SHOW TABLES")
	assert.NoError(t, err)
	for rows.Next() {
		result := make(map[string]interface{})
		err = rows.MapScan(result)
		fmt.Println(result)
	}
}
