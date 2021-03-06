package iter

import (
	"github.com/wwalexander/gofunc/opt"
)

type SliceIterator[T any] struct {
	slice []T
}

func FromSlice[T any](slice []T) *SliceIterator[T] {
	return &SliceIterator[T]{slice}
}

func (iter *SliceIterator[T]) Next() opt.Opt[T] {
	if len(iter.slice) == 0 {
		return opt.None[T]()
	}
	opt := opt.Some[T](iter.slice[0])
	iter.slice = iter.slice[1:]
	return opt
}

func ToSlice[T any, I Iter[T]](iter I) []T {
	return Reduce(iter, make([]T, 0), func(slice []T, v T) []T {
		return append(slice, v)
	})
}
