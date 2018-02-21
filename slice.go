package goaround

// ByteAt indexes string with bound safety.
func ByteAt(s string, index int) (ch byte) {
	if index < len(s) {
		ch = s[index]
		return ch
	} else {
		panic("string bounds out of range")
	}
}

// BoolAt indexes slices of booleans with bound safety.
func BoolAt(bs []bool, index int) (b bool) {
	if index < len(bs) {
		b = bs[index]
		return b
	} else {
		panic("slice bounds out of range")
	}
}

// StringAt indexes slices of strings with bound safety.
func StringAt(ss []string, index int) (s string) {
	if index < len(ss) {
		s = ss[index]
		return s
	} else {
		panic("slice bounds out of range")
	}
}

// IntAt indexes slices of integers with bound safety.
func IntAt(is []int, index int) (i int) {
	if index < len(is) {
		i = is[index]
		return i
	} else {
		panic("slice bounds out of range")
	}
}
