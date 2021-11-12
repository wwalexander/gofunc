package opt

// An Opt is an optional value of type T. An Opt can have some value of type T
// or have no value.
type Opt[T any] struct {
	v  T
	ok bool
}

// Some returns an Opt with value v.
func Some[T any](v T) Opt[T] {
	return Opt[T]{
		v:  v,
		ok: true,
	}
}

// None returns an Opt with no value.
func None[T any]() Opt[T] {
	return Opt[T]{}
}

// Let returns o's value. If o does not have a value, ok is false.
func (o Opt[T]) Let() (v T, ok bool) {
	if o.ok {
		v, ok = o.v, true
	}
	return
}

// Then returns an Opt containing the result of applying f to o's value.
// If o does not have a value, Then returns None.
func Then[T any, U any](o Opt[T], f func(T) U) Opt[U] {
	v, some := o.Let()
	if !some {
		return None[U]()
	}
	return Some(f(v))
}
