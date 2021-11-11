package opt

type Opt[T any] struct {
	v  T
	ok bool
}

func Some[T any](v T) Opt[T] {
	return Opt[T]{
		v:  v,
		ok: true,
	}
}

func None[T any]() Opt[T] {
	return Opt[T]{}
}

func (o Opt[T]) Let() (v T, ok bool) {
	if o.ok {
		v, ok = o.v, true
	}
	return
}

func Then[T any, U any](o Opt[T], f func(T) U) Opt[U] {
	v, some := o.Let()
	if !some {
		return None[U]()
	}
	return Some(f(v))
}
