package goaround

import (
	"fmt"
)


func ExampleByteAt() {
	fmt.Println(ByteAt("goaround", 2))
	// Output: 97
}
func ExampleByteAtExceeding() {
	fmt.Println(ByteAt("goaround", 100))
	// Output: 100
}

func ExampleBoolAt() {
	fmt.Println(BoolAt([]bool{true, false}, 0))
	// Output: true
}
func ExampleBoolAtExceeding() {
	fmt.Println(BoolAt([]bool{true, false}, 100))
	// Output: false
}

func ExampleIntAt() {
	fmt.Println(IntAt([]int{1, 2, 3}, 0))
	// Output: 1
}
func ExampleIntAtExceeding() {
	fmt.Println(IntAt([]int{1, 2, 3}, 100))
	// Output: 3
}

func ExampleStringAt() {
	fmt.Println(StringAt([]string{"go", "around"}, 0))
	// Output: go
}
func ExampleStringAtExceeding() {
	fmt.Println(StringAt([]string{"go", "around"}, 100))
	// Output: around
}
