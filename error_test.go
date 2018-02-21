package goaround

func ExampleFatalIf() {
	var err error
	FatalIf(err)
}

func ExampleLogIf() {
	var err error
	LogIf(err)
}

func ExamplePanicIf() {
	var err error
	PanicIf(err)
}
