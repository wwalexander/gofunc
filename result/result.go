package result

type Result[T any] struct {
	v   T
	err error
}

func Success[T any](v T) Result[T] {
	return Result[T]{
		v: v,
	}
}

func Failure[T any](err error) Result[T] {
	return Result[T]{
		err: err,
	}
}

func (r Result[T]) Let() (v T, err error) {
	if r.err == nil {
		v = r.v
	} else {
		err = r.err
	}
	return
}

func Then[T any, U any](r Result[T], f func(T) U) Result[U] {
	v, err := r.Let()
	if err != nil {
		return Failure[U](err)
	}
	return Success[U](f(v))
}
