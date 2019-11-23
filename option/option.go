package option

type Option struct {
	IsNone bool
	IsSome bool
	Value  interface{}
}

var none = Option{IsNone: true, IsSome: false}

func None() Option {
	return none
}

func Some(v interface{}) Option {
	return Option{IsNone: false, IsSome: true, Value: v}
}
