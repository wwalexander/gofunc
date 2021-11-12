package iter

import (
	"github.com/wwalexander/gofunc/opt"
)

// A FilterIterator is an Iter over each element of an Iter of type T that
// satisfies a predicate.
type FilterIterator[T any, I Iter[T]] struct {
	iter  I
	where func(T) bool
}

// Filter returns a FilterIterator over iter using the predicate where.
func Filter[T any, I Iter[T]](iter I, where func(T) bool) *FilterIterator[T, I] {
	return &FilterIterator[T, I]{
		iter:  iter,
		where: where,
	}
}

// Next returns the next element in iter.
func (iter *FilterIterator[T, I]) Next() opt.Opt[T] {
	return Find(iter.iter, iter.where)
}
