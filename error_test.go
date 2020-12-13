package goaround

func ExampleErrPrint() {
	ErrPrint("error message")
}

func ExampleExpect() {
	var err error
	Expect(err, "error message")
}
