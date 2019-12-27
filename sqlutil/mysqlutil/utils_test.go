package mysqlutil

import (
	"database/sql/driver"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/leizongmin/go/sqlutil"
)

func TestInterpolateParams(t *testing.T) {
	args := []driver.Value{"tt", 123, "xxx", time.Now()}
	ret, err := sqlutil.InterpolateParams("SELECT * FROM ?? WHERE a=? AND b=? AND c=?", args, time.Local, escapeStringQuotes)
	assert.NoError(t, err)
	fmt.Println(ret)
}

func TestInterpolateParams2(t *testing.T) {
	args := []driver.Value{`\\"da\\"`}
	ret, err := sqlutil.InterpolateParams("VALUES (?)", args, time.Local, escapeStringQuotes)
	if err != nil {
		t.Error(err)
	}
	assert.NoError(t, err)
	assert.Equal(t, `VALUES ('\\\\"da\\\\"')`, ret)
	fmt.Println(ret)
}
