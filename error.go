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

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrPrint(message string) {
	_, _ = os.Stderr.WriteString(message)
}
