package traineer

// sliceContains checks if a slice of strings contains a certain string.
func sliceContains(slc []string, str string) bool {
	for _, s := range slc {
		if s == str {
			return true
		}
	}

	return false
}
