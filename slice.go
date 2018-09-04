package goaround

func normalizeIndex(index int, length int) int {
	if index < 0 {
		var negatived int = length + index
		if negatived < 0 {
			return 0
		} else {
			return negatived
		}
	} else if index >= length {
		return length - 1
	} else{
		return index
	}
}

// GetByte indexes string with bound safety.
func GetByte(s string, index int) byte {
	var length int = len(s)
	var normalizedIndex int = normalizeIndex(index, length)
	return s[normalizedIndex]
}

// GetBool indexes slices of booleans with bound safety.
func GetBool(s []bool, index int) bool {
	var length int = len(s)
	var normalizedIndex int = normalizeIndex(index, length)
	return s[normalizedIndex]
}

// GetString indexes slices of strings with bound safety.
func GetString(s []string, index int) string {
	var length int = len(s)
	var normalizedIndex int = normalizeIndex(index, length)
	return s[normalizedIndex]
}

// GetInt indexes slices of integers with bound safety.
func GetInt(s []int, index int) int {
	var length int = len(s)
	var normalizedIndex int = normalizeIndex(index, length)
	return s[normalizedIndex]
}

	} else {
		return intSlice[index]
	}
}
