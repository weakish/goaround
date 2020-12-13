package goaround

func ExamplePanicIf() {
	var err error
	PanicIf(err)
}

func ExampleErrPrint() {
	ErrPrint("error message")
}

func ExampleExpect() {
	var err error
	Expect(err, "error message")
}
