package typeutil

import (
	"fmt"
	"testing"
)

func TestAny(t *testing.T) {
	var a Any
	a = 123
	fmt.Println(a)
	a = "xxx"
	fmt.Println(a)
}
