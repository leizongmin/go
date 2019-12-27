package configloader

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"

	"github.com/leizongmin/go/randutil"
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

func TestLoadFile(t *testing.T) {
	f := filepath.Join(os.TempDir(), fmt.Sprintf("%f.json", randutil.Float64()))
	err := ioutil.WriteFile(f, []byte(`{"a":123,"b":"xxx"}`), 0644)
	assert.NoError(t, err)

	v := struct {
		A int    `json:"a"`
		B string `json:"b"`
	}{}
	err = LoadFile(f, &v)
	assert.NoError(t, err)
	assert.Equal(t, 123, v.A)
	assert.Equal(t, "xxx", v.B)
	litter.Dump(v)

	err = LoadFile(f[:len(f)-5]+".*", &v)
	assert.NoError(t, err)
	assert.Equal(t, 123, v.A)
	assert.Equal(t, "xxx", v.B)
	litter.Dump(v)
}
