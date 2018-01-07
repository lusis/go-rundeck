package testdata

// GetBytes returns a test file data as bytes. Useful for testing
func GetBytes(file string) ([]byte, error) {
	return Asset(file)
}
