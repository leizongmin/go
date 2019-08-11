package configLoader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	v := struct {
		A int    `json:"a"`
		B string `json:"b"`
	}{}
	err := Load("json", []byte(`{"a":123,"b":"xxx"}`), &v)
	assert.NoError(t, err)
	assert.Equal(t, 123, v.A)
	assert.Equal(t, "xxx", v.B)
}
