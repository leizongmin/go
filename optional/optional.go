package optional

type Value struct {
	value interface{}
	none  bool
}

var none = Value{none: true}

// Value值
func Some(v interface{}) Value {
	return Value{value: v}
}

// None值
func None() Value {
	return none
}

// 是否为None
func (o Value) IsNone() bool {
	return o.none
}

// 是否为Value
func (o Value) IsValue() bool {
	return !o.none
}

// 获得Value
func (o Value) Value() interface{} {
	return o.value
}
