package goaround

// GetByte indexes string with bound safety.
func GetByte(someString string, index int) byte {
	var length int = len(someString)
	if index < 0 {
		var negatived int = length + index
		if negatived < 0 {
			return someString[0]
		} else {
			return someString[negatived]
		}
	} else if index >= length {
		return someString[length - 1]
	} else {
		return someString[index]
	}
}

// GetBool indexes slices of booleans with bound safety.
func GetBool(booleanSlice []bool, index int) bool {
	var length int = len(booleanSlice)
	if index < 0 {
		var negatived int = length + index
		if negatived < 0 {
			return booleanSlice[0]
		} else {
			return booleanSlice[negatived]
		}

	} else if index >= length {
		return booleanSlice[length - 1]
	} else {
		return booleanSlice[index]
	}
}

// GetString indexes slices of strings with bound safety.
func GetString(stringSlice []string, index int) string {
	var length int = len(stringSlice)
	if index < 0 {
		var negatived int = length + index
		if negatived < 0 {
			return stringSlice[0]
		} else {
			return stringSlice[negatived]
		}
	} else if index >= length {
		return stringSlice[length - 1]
	} else {
		return stringSlice[index]
	}
}

// GetInt indexes slices of integers with bound safety.
func GetInt(intSlice []int, index int) int {
	var length int = len(intSlice)
	if index < 0 {
		var negatived int = length + index
		if negatived < 0 {
			return intSlice[0]
		} else {
			return intSlice[negatived]
		}
	} else if index >= length {
		return intSlice[length - 1]
	} else {
		return intSlice[index]
	}
}
