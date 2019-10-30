// Package goaround defines some workarounds for golang.
//
// - Check nullability for function, interface, and pointers.
// - Bound safe index methods of common types of slices.

package goaround

// RequireNonNull will panic if input is nil.
func RequireNonNull(any interface{}) {
	if any == nil {
		panic("must not be null")
	}
}
