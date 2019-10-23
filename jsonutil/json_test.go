package jsonutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustMarshal(t *testing.T) {
	assert.Equal(t, string(MustMarshal(&map[string]interface{}{
		"a": 123,
		"b": "xxx",
	})), `{"a":123,"b":"xxx"}`)
}

func TestMustUnmarshal(t *testing.T) {
	v := map[string]interface{}{}
	MustUnmarshal([]byte(`{"a":123,"b":"xxx"}`), &v)
	assert.Equal(t, float64(123), v["a"])
	assert.Equal(t, "xxx", v["b"])
}
