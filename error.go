package goaround

import (
	"os"
)

// PanicIf is equivalent to, but more explict than, using a blank identifier.
func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrPrint(message string) {
	_, _ = os.Stderr.WriteString(message)
}

// Inspired by Rust.
func Expect(err error, message string) {
	if err != nil {
		ErrPrint(message)
		panic(err)
	}
}
