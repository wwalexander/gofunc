package iter

import (
	"github.com/wwalexander/gofunc/opt"
)

// An Iter returns elements of type T from an underlying sequence.
type Iter[T any] interface {
	// Next returns the next element in the sequence.
	// If there are no more elements in the sequence, Next returns None.
	Next() opt.Opt[T]
}

// ForEach applies f to each element in iter.
func ForEach[T any, I Iter[T]](iter I, f func(T)) {
	for {
		elem, some := iter.Next().Let()
		if !some {
			break
		}
		f(elem)
	}
}

// Find finds the first element in iter satisfying the predicate where.
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

// Reduce reduces the elements of iter into a single value of type A.
// Reduce accumulates the result of applying f to the running accumulated value
// and each element in iter, starting from initial.
func Reduce[T any, A any, I Iter[T]](iter I, initial A, f func(A, T) A) A {
	reduced := initial
	ForEach(iter, func(v T) {
		reduced = f(reduced, v)
	})
	return reduced
}
