package sqlBuilder

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInterpolateParams(t *testing.T) {
	args := []Value{"tt", 123, "xxx", time.Now()}
	ret, err := InterpolateParams("SELECT * FROM ?? WHERE a=? AND b=? AND c=?", args, time.Local)
	assert.NoError(t, err)
	fmt.Println(ret)
}

func TestInterpolateParams2(t *testing.T) {
	args := []Value{`\\"da\\"`}
	ret, err := InterpolateParams("VALUES (?)", args, time.Local)
	if err != nil {
		t.Error(err)
	}
	assert.NoError(t, err)
	assert.Equal(t, `VALUES ('\\\\"da\\\\"')`, ret)
	fmt.Println(ret)
}
