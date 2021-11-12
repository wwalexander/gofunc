package iter

import (
	"github.com/wwalexander/gofunc/opt"
)

type ChunkIterator[T any, I Iter[T]] struct {
	iter I
	size int
}

func Chunk[T any, I Iter[T]](iter I, size int) *ChunkIterator[T, I] {
	return &ChunkIterator[T, I]{
		iter: iter,
		size: size,
	}
}

func (iter *ChunkIterator[T, I]) Next() opt.Opt[[]T] {
	var chunk []T
	for i := 0; i < iter.size; i++ {
		elem, ok := iter.iter.Next().Let()
		if !ok {
			break
		}
		chunk = append(chunk, elem)
	}
	if chunk == nil {
		return opt.None[[]T]()
	}
	return opt.Some(chunk)
}
