package iter

import (
	"github.com/wwalexander/gofunc/opt"
)

// A MapIterator is an Iter which transforms each element of an Iter of type T
// to elements of type U.
type MapIterator[T any, U any, I Iter[T]] struct {
	iter I
	f    func(T) U
}

// Map returns a MapIterator transforming the elements of iter by applying f.
func Map[T any, U any, I Iter[T]](iter I, f func(T) U) *MapIterator[T, U, I] {
	return &MapIterator[T, U, I]{
		iter: iter,
		f:    f,
	}
}

// Next returns the next element in iter.
func (iter *MapIterator[T, U, I]) Next() opt.Opt[U] {
	next := iter.iter.Next()
	return opt.Then(next, iter.f)
}
