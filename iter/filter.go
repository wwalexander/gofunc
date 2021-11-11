package iter

import (
	"github.com/wwalexander/gofunc/opt"
)

type FilterIterator[T any, I Iter[T]] struct {
	iter  I
	where func(T) bool
}

func Filter[T any, I Iter[T]](iter I, where func(T) bool) *FilterIterator[T, I] {
	return &FilterIterator[T, I]{
		iter:  iter,
		where: where,
	}
}

func (iter *FilterIterator[T, I]) Next() opt.Opt[T] {
	return Find(iter.iter, iter.where)
}
