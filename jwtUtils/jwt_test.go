package jwtUtils

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	jwt := New("hello")
	str, err := jwt.Sign(Value{
		"a": 123,
		"b": "aaa",
	})
	if err != nil {
		t.Errorf("Sign: %s", err)
	}
	fmt.Println(str)
	data, err := jwt.UnSign(str)
	if err != nil {
		t.Errorf("UnSign: %s", err)
	}
	fmt.Printf("%+v\n", data)
	if data["a"].(float64) != 123 {
		t.Fail()
	}
	if data["b"].(string) != "aaa" {
		t.Fail()
	}
}
