package optional

type Optional struct {
	IsNone bool
	IsSome bool
	Value  interface{}
}

var none = Optional{IsNone: true, IsSome: false}

func None() Optional {
	return none
}

func Some(v interface{}) Optional {
	return Optional{IsNone: false, IsSome: true, Value: v}
}
