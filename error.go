package goaround

import (
	"log"
	"os"
)

func FatalIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogIf(err error) {
	if err != nil {
		log.Print(err)
	}
}

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