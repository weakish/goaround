package goaround

// GetByte indexes string with bound safety.
func GetByte(someString string, index int) byte {
	var length int = len(someString)
	if index < length {
		return someString[index]
	} else {
		return someString[length - 1]
	}
}

// GetBool indexes slices of booleans with bound safety.
func GetBool(booleanSlice []bool, index int) bool {
	var length int = len(booleanSlice)
	if index < length {
		return booleanSlice[index]
	} else {
		return booleanSlice[length - 1]
	}
}

// GetString indexes slices of strings with bound safety.
func GetString(stringSlice []string, index int) string {
	var length int = len(stringSlice)
	if index < length {
		return stringSlice[index]
	} else {
		return stringSlice[length - 1]
	}
}

// GetInt indexes slices of integers with bound safety.
func GetInt(intSlice []int, index int) int {
	var length int = len(intSlice)
	if index < length {
		return intSlice[index]
	} else {
		return intSlice[length - 1]
	}
}
