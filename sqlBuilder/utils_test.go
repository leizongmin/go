package sqlBuilder

import (
	"testing"
	"time"
)

func TestInterpolateParams(t *testing.T) {
	args := []Value{"tt", 123, "xxx", time.Now()}
	ret, err := InterpolateParams("SELECT * FROM ?? WHERE a=? AND b=? AND c=?", args, time.Local)
	if err != nil {
		t.Error(err)
	}
	println(ret)
}
