package goaround

import "testing"

func TestRequireNonNullShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("RequireNonNull should panic!")
		}
	}()

	RequireNonNull(nil)
}

func TestRequireNonNullNotPanic(t *testing.T) {
	RequireNonNull(0)
}

