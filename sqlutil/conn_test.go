package sqlutil

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestFindMany(t *testing.T) {
	EnableDebug()
	db, err := Open("mysql", "root:@tcp(localhost:3306)/mysql")
	assert.NoError(t, err)
	fmt.Printf("%+v\n", db)

	count, ok := QueryCount(db, "SELECT COUNT(*) AS count FROM user")
	fmt.Println(count, ok)

	var list []Row
	var tx AbstractTx
	tx = db.MustBegin()

	ok = QueryMany(tx, &list, "SHOW TABLES")
	fmt.Printf("%v %+v\n", ok, list)

	rows, err := db.Queryx("SHOW TABLES")
	assert.NoError(t, err)
	for rows.Next() {
		result := make(map[string]interface{})
		err = rows.MapScan(result)
		fmt.Println(result)
	}

	list = []Row{}
	ok = QueryMany(tx, &list, "SHOW TABLES")
	fmt.Printf("%v %+v\n", ok, list)

	tx2, err := db.Beginx()
	assert.NoError(t, err)
	ok = QueryMany(tx2, &list, "SHOW TABLES")
	fmt.Printf("%v %+v\n", ok, list)

	var db2 AbstractDB
	db2 = db
	fmt.Println(db2)
	var db3 AbstractDBBase
	db3 = db
	fmt.Println(db3)
	var tx3 AbstractTx
	tx3 = tx2
	fmt.Println(tx3)

	{
		user, ok := QueryOneToMap(db, "SELECT count(*) as count FROM user LIMIT 1")
		assert.True(t, ok)
		fmt.Println(user)
		for n, v := range user {
			fmt.Println("\t", n, string(v.([]byte)))
		}

		users, ok := QueryManyToMap(db, "SELECT * FROM user LIMIT 3")
		assert.True(t, ok)
		fmt.Println(users)
		for _, user := range users {
			fmt.Println(user)
			for n, v := range user {
				if v != nil {
					fmt.Println("\t", n, string(v.([]byte)))
				}
			}
		}
	}

	{
		err := BatchExec(db, "SELECT * FROM user LIMIT 1", "SELECT * FROM user LIMIT 1")
		assert.NoError(t, err)
	}
}
