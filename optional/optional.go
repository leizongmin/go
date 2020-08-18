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

func (o Optional) Map(mapper func(value interface{}) interface{}) Optional {
	if o.IsNone {
		return o
	}
	return Some(mapper(o.Value))
}

func (o Optional) FlatMap(mapper func(value interface{}) Optional) Optional {
	if o.IsNone {
		return o
	}
	return mapper(o.Value)
}

func (o Optional) Filter(predicate func(value interface{}) bool) Optional {
	if o.IsNone {
		return o
	}
	if predicate(o.Value) {
		return o
	}
	return none
}

func (o Optional) OrElse(otherValue interface{}) interface{} {
	if o.IsNone {
		return otherValue
	}
	return o.Value
}

func (o Optional) OrElseGet(supplier func() interface{}) interface{} {
	if o.IsNone {
		return supplier()
	}
	return o.Value
}
