package goaround

// Inspired by Rust.
func Expect(err error, message string) {
	if err != nil {
		println(message)
		panic(err)
	}
}
