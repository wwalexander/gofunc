package iter

import (
	"github.com/wwalexander/gofunc/opt"
)

type MapIterator[T any, U any, I Iter[T]] struct {
	iter I
	f    func(T) U
}

func Map[T any, U any, I Iter[T]](iter I, f func(T) U) *MapIterator[T, U, I] {
	return &MapIterator[T, U, I]{
		iter: iter,
		f:    f,
	}
}

func (iter *MapIterator[T, U, I]) Next() opt.Opt[U] {
	next := iter.iter.Next()
	return opt.Then(next, iter.f)
}
