package tryRecover

import (
	"testing"
)

func TestEat(t *testing.T) {
	defer Eat()
	panic("Some error")
}
