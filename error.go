package goaround

import "log"

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
