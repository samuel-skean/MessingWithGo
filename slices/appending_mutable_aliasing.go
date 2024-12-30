package main

import "fmt"

func appendSomeAndShow(x, y []int) {
	fmt.Println(x, y)
	fmt.Println(cap(x), cap(y))
	y = append(y, 20)
	fmt.Println(x, y)
}

func main() {
	x := []int{}
	fmt.Println(x)
	x = append(x, 1, 2, 3)

	// Third argument to slicing expression sets maximum index accessibly through capacity, not the step:
	// (https://go.dev/ref/spec#Slice_expressions)
	appendSomeAndShow(x, x[:0:0])

	// Otherwise, defaults to maximum capacity (or length, I haven't checked) of original slice, and appending can
	// overwrite in the original slice. Eek!
	appendSomeAndShow(x, x[:0])

	// Setting the capacity to high panics.
	_ = x[:0:20]
}
