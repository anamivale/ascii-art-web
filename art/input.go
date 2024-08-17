package art

// CheckInputString checks if the input string is between the provided range.
func CheckInputString(input string) bool {
	for _, char := range input {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
