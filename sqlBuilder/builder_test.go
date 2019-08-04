package sqlBuilder

import (
	"fmt"
	"testing"
)

func TestTable(t *testing.T) {
	sql := Table("test").Select("*").Where("a=?", 123).And("b=?", 456).Skip(10).Limit(20).Build()
	fmt.Println(sql)
}
