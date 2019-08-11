package randUtils

import (
	"testing"
	"time"
)

func TestSetSeed(t *testing.T) {
	SetSeed(time.Now().Unix())
	Int63n(123)
}
