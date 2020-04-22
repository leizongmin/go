package zaputil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	log, err := Create(nil)
	defer log.Sync()
	assert.NoError(t, err)
	log.Info("hello")
}
