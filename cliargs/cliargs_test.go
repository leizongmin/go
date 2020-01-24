package cliargs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	a := Parse([]string{"x=123", "--a=123", "-b=456", "c", "def"})
	fmt.Printf("%+v\n", a)

	assert.Equal(t, 2, a.OptionsCount())
	assert.True(t, a.HasOption("a"))
	assert.True(t, a.HasOption("b"))
	assert.False(t, a.HasOption("x"))
	assert.False(t, a.HasOption("c"))
	assert.False(t, a.HasOption("def"))

	assert.Equal(t, 3, a.ArgsCount())
	assert.Equal(t, "x=123", a.GetArg(0))
	assert.Equal(t, "c", a.GetArg(1))
	assert.Equal(t, "def", a.GetArg(2))

	b := a.SubArgs(1)

	assert.Equal(t, 2, a.OptionsCount())
	assert.True(t, a.HasOption("a"))
	assert.True(t, a.HasOption("b"))
	assert.False(t, a.HasOption("x"))
	assert.False(t, a.HasOption("c"))
	assert.False(t, a.HasOption("def"))

	assert.Equal(t, 2, b.ArgsCount())
	assert.Equal(t, "c", b.GetArg(0))
	assert.Equal(t, "def", b.GetArg(1))

	a.ForEachArgs(func(item string) {
		fmt.Println(item)
	})
	a.ForEachOptions(func(item OptionItem) {
		fmt.Println(item)
	})
}
