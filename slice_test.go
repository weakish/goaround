package goaround

import (
	"testing"
	"fmt"
)

func TestByteAt(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Index out of range should panic!")
		}
	}()

	ByteAt("", 128)
}

func ExampleByteAt() {
	fmt.Println(ByteAt("goaround", 2))
	// Output: 97
}

func ExampleBoolAt() {
	fmt.Println(BoolAt([]bool{true, false}, 0))
	// Output: true
}

func ExampleIntAt() {
	fmt.Println(IntAt([]int{1, 2, 3}, 0))
	// Output: 1
}

func ExampleStringAt() {
	fmt.Println(StringAt([]string{"go", "around"}, 0))
	// Output: go
}
