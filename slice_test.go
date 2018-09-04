package goaround

import (
	"fmt"
)


func ExampleGetByte() {
	fmt.Println(GetByte("goaround", 2))
	// Output: 97
}
func ExampleGetByteExceeding() {
	fmt.Println(GetByte("goaround", 100))
	// Output: 100
}
func ExampleGetByteNegative() {
	fmt.Println(GetByte("goaround", -1))
	// Output: 100
}
func ExampleGetByteLeftExceeding() {
	fmt.Println(GetByte("goaround", -100))
	// Output: 103
}

func ExampleGetBool() {
	fmt.Println(GetBool([]bool{true, false}, 0))
	// Output: true
}
func ExampleGetBoolExceeding() {
	fmt.Println(GetBool([]bool{true, false}, 100))
	// Output: false
}
func ExampleGetBoolNegative() {
	fmt.Println(GetBool([]bool{true, false}, -1))
	// Output: false
}
func ExampleGetBoolLeftExceeding() {
	fmt.Println(GetBool([]bool{true, false}, -100))
	// Output: true
}

func ExampleGetInt() {
	fmt.Println(GetInt([]int{1, 2, 3}, 0))
	// Output: 1
}
func ExampleGetIntExceeding() {
	fmt.Println(GetInt([]int{1, 2, 3}, 100))
	// Output: 3
}
func ExampleGetIntNegative() {
	fmt.Println(GetInt([]int{1, 2, 3}, -1))
	// Output: 3
}
func ExampleGetIntLeftExceeding() {
	fmt.Println(GetInt([]int{1, 2, 3}, -100))
	// Output: 1
}

func ExampleGetString() {
	fmt.Println(GetString([]string{"go", "around"}, 0))
	// Output: go
}
func ExampleGetStringExceeding() {
	fmt.Println(GetString([]string{"go", "around"}, 100))
	// Output: around
}
func ExampleGetStringNegative() {
	fmt.Println(GetString([]string{"go", "around"}, -1))
	// Output: around
}
func ExampleGetStringLeftExceeding() {
	fmt.Println(GetString([]string{"go", "around"}, -100))
	// Output: go
}
