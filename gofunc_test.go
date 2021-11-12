package gofunc_test

import (
	"fmt"
	"math"
	"strings"

	"github.com/wwalexander/gofunc/iter"
	"github.com/wwalexander/gofunc/opt"
)

func ExampleOpt() {
	hackers := []struct {
		Name   string
		Handle opt.Opt[string]
	}{
		{"Dade Murphy", opt.Some("Crash Override")},
		{"Kate Libby", opt.Some("Acid Burn")},
		{"Joey Pardella", opt.None[string]()},
	}
	for _, hacker := range hackers {
		fmt.Printf("%s: ", hacker.Name)
		identity := opt.Then(hacker.Handle, func(handle string) string {
			return fmt.Sprintf("I'm %s.", handle)
		})
		if introduction, ok := identity.Let(); ok {
			fmt.Println(introduction)
		} else {
			fmt.Println("I don't have an identity until I have a handle.")
		}
	}
	// Output:
	// Dade Murphy: I'm Crash Override.
	// Kate Libby: I'm Acid Burn.
	// Joey Pardella: I don't have an identity until I have a handle.
}

func ExampleForEach() {
	designers := iter.FromSlice([]string{
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
	designers := iter.FromSlice([]string{
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
	designers := iter.FromSlice([]string{
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
