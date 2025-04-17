package validator

// IsInRange checks if a number is within the specified range (inclusive)
func IsInRange(num, min, max int) bool {
	// Check if the number is less than the minimum
	if num < min {
		return false
	}

	// Check if the number is greater than the maximum
	if num > max {
		return false
	}

	// If we get here, the number is within range
	return true
}
