package sqlutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	{
		sql := Table("test").Select().Where("a=?", 123).And("b=?", 456).Skip(10).Limit(20).Build()
		fmt.Println(sql)
		assert.Equal(t, "SELECT * FROM test WHERE a=123 AND b=456 LIMIT 10,20", sql)
	}
	{
		sql := Table("test").Select().WhereRow(Row{"a": 123, "b": 456}).Skip(10).Limit(20).Build()
		fmt.Println(sql)
		assert.Equal(t, "SELECT * FROM test WHERE a=123 AND b=456 LIMIT 10,20", sql)
	}
	{
		sql := Table("test").Select("*").Where("a=?", 123).And("b=?", 456).Skip(10).Limit(20).Build()
		fmt.Println(sql)
		assert.Equal(t, "SELECT * FROM test WHERE a=123 AND b=456 LIMIT 10,20", sql)
	}
	{
		sql := Table("test").Select("*").Where("a=?", 123).And("b=?", `'xxx'`).Skip(10).Limit(20).Build()
		fmt.Println(sql)
		assert.Equal(t, "SELECT * FROM test WHERE a=123 AND b='''xxx''' LIMIT 10,20", sql)
	}
	{
		sql := Table("test").Insert(Row{
			"a": 123,
			"b": `'ok'`,
			"c": true,
		}).Build()
		fmt.Println(sql)
		assert.Equal(t, "INSERT INTO test (a, b, c) VALUES (123, '''ok''', TRUE)", sql)
	}
	{
		sql := Table("test").InsertMany([]Row{
			{
				"a": 123,
				"b": `'ok'`,
				"c": true,
			},
			{
				"c": 666,
				"b": "current_timestamp()",
			},
		}).Build()
		fmt.Println(sql)
		assert.Equal(t, "INSERT INTO test (a, b, c) VALUES (123, '''ok''', TRUE), (NULL, 'current_timestamp()', 666)", sql)
	}
	{
		sql := Table("test").Insert(Row{
			"a": 123,
			"b": `'ok'`,
			"c": true,
		}).OnDuplicateKeyUpdate().Build()
		fmt.Println(sql)
		assert.Equal(t, "INSERT INTO test (a, b, c) VALUES (123, '''ok''', TRUE) ON DUPLICATE KEY UPDATE", sql)
	}
	{
		sql := Table("test").Insert(Row{
			"a": 123,
			"b": `'ok'`,
			"c": true,
		}).ReturningAll().Build()
		fmt.Println(sql)
		assert.Equal(t, "INSERT INTO test (a, b, c) VALUES (123, '''ok''', TRUE) RETURNING *", sql)
	}
	{
		sql := Table("test").Insert(Row{
			"a": 123,
			"b": `'ok'`,
			"c": true,
		}).Returning("a", "b", "c").Build()
		fmt.Println(sql)
		assert.Equal(t, "INSERT INTO test (a, b, c) VALUES (123, '''ok''', TRUE) RETURNING a, b, c", sql)
	}
	{
		sql := Table("test").Count("*").Build()
		fmt.Println(sql)
		assert.Equal(t, "SELECT COUNT(*) AS count FROM test", sql)
	}
	{
		sql := Table("test").Delete().Where("a=?", true).Limit(1).Build()
		fmt.Println(sql)
		assert.Equal(t, "DELETE FROM test WHERE a=TRUE LIMIT 1", sql)
	}
	{
		sql := Table("test").Update().Set("a=?, b=?", 123, 456).Set("c=now()").Where("a=?", 999).Limit(10).Build()
		fmt.Println(sql)
		assert.Equal(t, "UPDATE test SET a=123, b=456, c=now() WHERE a=999 LIMIT 10", sql)
	}
	{
		sql := Table("test").Update().SetRow(Row{
			"a": 123,
			"b": 456,
			"c": "xxx",
		}).Where("a=?", 999).Limit(10).Build()
		fmt.Println(sql)
		assert.Equal(t, "UPDATE test SET a=123, b=456, c='xxx' WHERE a=999 LIMIT 10", sql)
	}
	{
		sql := Table("test").Update().SetRow(Row{
			"a": 123,
			"b": 456,
			"c": "xxx",
		}).Where("a=?", 999).ReturningAll().Build()
		fmt.Println(sql)
		assert.Equal(t, "UPDATE test SET a=123, b=456, c='xxx' WHERE a=999 RETURNING *", sql)
	}
	{
		sql := Table("test").Select("a", "b").LeftJoin("test2", "c").Build()
		fmt.Println(sql)
		assert.Equal(t, "SELECT a, b, c FROM test LEFT JOIN test2", sql)
	}
	{
		sql := Table("test").Select("a", "b").As("x").LeftJoin("test2", "c").As("y").On("x.a=y.a").RightJoin("test3").As("z").Where("x.a=666").Build()
		fmt.Println(sql)
		assert.Equal(t, "SELECT a, b, c FROM test AS x LEFT JOIN test2 AS y ON x.a=y.a RIGHT JOIN test3 AS z WHERE x.a=666", sql)
	}
	{
		sql := Table("test").SelectDistinct("*").Where("a=?", 123).And("b=?", 456).Skip(10).Limit(20).Build()
		fmt.Println(sql)
		assert.Equal(t, "SELECT DISTINCT * FROM test WHERE a=123 AND b=456 LIMIT 10,20", sql)
	}
	{
		sql := Table("test").SelectDistinct("*").Where("a=?", 123).And("b=?", 456).GroupBy("a").Having("a=b").OrderBy("b").Skip(10).Limit(20).Build()
		fmt.Println(sql)
		assert.Equal(t, "SELECT DISTINCT * FROM test WHERE a=123 AND b=456 GROUP BY a HAVING a=b ORDER BY b LIMIT 10,20", sql)
	}
	{
		q1 := Table("test").Select().As("a")
		q2 := q1.Clone().As("b")
		sql1 := q1.Build()
		sql2 := q2.Build()
		fmt.Println(sql1)
		fmt.Println(sql2)
		assert.Equal(t, "SELECT * FROM test AS a", sql1)
		assert.Equal(t, "SELECT * FROM test AS b", sql2)
	}
}
