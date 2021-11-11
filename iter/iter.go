package iter

import (
	"github.com/wwalexander/gofunc/opt"
)

type Iter[T any] interface {
	Next() opt.Opt[T]
}

func ForEach[T any, I Iter[T]](iter I, f func(T)) {
	for {
		elem, some := iter.Next().Let()
		if !some {
			break
		}
		f(elem)
	}
}

func Find[T any, I Iter[T]](iter I, where func(T) bool) opt.Opt[T] {
	for {
		v, ok := iter.Next().Let()
		if !ok {
			break
		}
		if where(v) {
			return opt.Some[T](v)
		}
	}
	return opt.None[T]()
}

func Reduce[T any, A any, I Iter[T]](iter I, initial A, f func(A, T) A) A {
	reduced := initial
	ForEach(iter, func(v T) {
		reduced = f(reduced, v)
	})
	return reduced
}
