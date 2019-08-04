package sqlBuilder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	{
		sql := Table("test").Select("*").Where("a=?", 123).And("b=?", 456).Skip(10).Limit(20).Build()
		fmt.Println(sql)
		assert.Equal(t, "SELECT * FROM test WHERE a=123 AND b=456  LIMIT 10,20", sql)
	}
	{
		sql := Table("test").Select("*").Where("a=?", 123).And("b=?", `'xxx'`).Skip(10).Limit(20).Build()
		fmt.Println(sql)
		assert.Equal(t, `SELECT * FROM test WHERE a=123 AND b='''xxx'''  LIMIT 10,20`, sql)
	}
	{
		sql := Table("test").Insert(map[string]interface{}{
			"a": 123,
			"b": `'ok'`,
			"c": true,
		}).Build()
		fmt.Println(sql)
		assert.Equal(t, `INSERT INTO test (b, c, a) VALUES ('''ok''', 1, 123)`, sql)
	}
}
