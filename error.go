package goaround

import (
	"os"
)

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
