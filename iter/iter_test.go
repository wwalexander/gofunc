package iter_test

import (
	"fmt"
	"math"
	"strings"

	"github.com/wwalexander/gofunc/iter"
)

func ExampleForEach() {
	designers := iter.FromSlice[string]([]string{
		"Robert Griesemer",
		"Rob Pike",
		"Ken Thompson",
	})
	iter.ForEach(designers, func(designer string) {
		fmt.Println(designer)
	})
	// Output:
	// Robert Griesemer
	// Rob Pike
	// Ken Thompson
}

func ExampleFind() {
	designers := iter.FromSlice[string]([]string{
		"Robert Griesemer",
		"Ken Thompson",
		"Rob Pike",
	})
	nonRob := iter.Find(designers, func(designer string) bool {
		return !strings.HasPrefix(designer, "Rob")
	})
	if designer, ok := nonRob.Let(); ok {
		fmt.Println(designer)
	}
	// Output: Ken Thompson
}

func ExampleReduce() {
	devs := iter.FromSlice([]int{
		3, 7, 2, 0, 1,
	})
	sumOfSquares := iter.Reduce(devs, 0, func(sum int, dev int) int {
		return sum + dev*dev
	})
	fmt.Println(sumOfSquares)
	// Output: 63
}

func ExampleChunk() {
	data := iter.FromSlice([]byte{
		255, 0, 0,
		0, 255, 0,
		0, 0, 255,
	})
	pixels := iter.Chunk[byte](data, 3)
	iter.ForEach(pixels, func(pixel []byte) {
		fmt.Printf("#%02X%02X%02X\n", pixel[0], pixel[1], pixel[2])
	})
	// Output:
	// #FF0000
	// #00FF00
	// #0000FF
}

func ExampleMap() {
	numbers := iter.FromSlice([]float64{
		2.2, 0.9, 3.5, 1.8,
	})
	rounded := iter.Map(numbers, func(number float64) int {
		return int(math.Round(number))
	})
	iter.ForEach(rounded, func(number int) {
		fmt.Println(number)
	})
	// Output:
	// 2
	// 1
	// 4
	// 2
}

func ExampleFilter() {
	designers := iter.FromSlice[string]([]string{
		"Robert Griesemer",
		"Ken Thompson",
		"Rob Pike",
	})
	robs := iter.Filter(designers, func(designer string) bool {
		return strings.HasPrefix(designer, "Rob")
	})
	iter.ForEach(robs, func(rob string) {
		fmt.Println(rob)
	})
	// Output:
	// Robert Griesemer
	// Rob Pike
}
