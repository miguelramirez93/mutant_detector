package arrayutils

// IsStringInRange check if strToCheck only contains elements in arrRange
func IsStringInRange(strToCheck string, arrRange []string) bool {

	for _, runeToCheck := range strToCheck {
		foundElements := 0
		for _, rangeElement := range arrRange {
			strElement := string(runeToCheck)
			if strElement == rangeElement {
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
