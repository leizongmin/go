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

	count, ok := FindCount(db, "SELECT COUNT(*) AS count FROM user")
	fmt.Println(count, ok)

	var list []Row
	var tx AbstractTx
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
