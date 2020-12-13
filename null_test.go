package goaround

import "testing"

func TestNotNil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("NotNil should panic!")
		}
	}()

	NotNil(nil)
}

func ExampleNotNil() {
	NotNil(0)
}
