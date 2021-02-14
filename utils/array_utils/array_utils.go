package arrayutils

// IsStringInRange check if strToCheck only contains elements in arrRange
func IsStringInRange(strToCheck string, arrRange []string) bool {
	foundElements := 0
	for _, rangeElement := range arrRange {
		for _, strElement := range strToCheck {
			if string(strElement) == rangeElement {
				foundElements++
				break
			}
		}
		if foundElements == 0 {
			return false
		}
	}
	return true
}
