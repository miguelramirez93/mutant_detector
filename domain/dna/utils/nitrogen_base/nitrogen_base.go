package nitrogenbaseutils

import "math"

// GetHorizontalCoincidences return nitrogen base horizontal coincidences in the given dna matrix
// If macCoincidences > 0 will stop after find maxCoincidences.
func GetHorizontalCoincidences(dna []string, repeatRange, maxCoincidences int) int {

	x1 := 0

	maxCoincidencesPerGroup := math.Floor(float64((len(dna[0])) / repeatRange))
	currentCoincidences := 0

	if maxCoincidencesPerGroup == 0 {
		return 0
	}

	fixedRepeatRange := repeatRange - 1

	for x1 < len(dna) {
		rowCoincidences := 0
		y1 := 0
		x2 := x1
		y2 := y1 + int(repeatRange)
		for y1 <= (len(dna[x1]) - repeatRange) {
			y2 = y1 + fixedRepeatRange
			nitrogenBaseCoincidences := 0
			for y2 > y1 {
				xy1 := string(dna[x1][y1])
				xy2 := string(dna[x2][y2])
				if xy1 == xy2 {
					y2--
					nitrogenBaseCoincidences++

				} else {
					break
				}
			}
			if nitrogenBaseCoincidences == fixedRepeatRange {
				rowCoincidences++
				currentCoincidences++
				if currentCoincidences == maxCoincidences {
					return currentCoincidences
				}

				if rowCoincidences == int(maxCoincidencesPerGroup) {
					break
				}

			}
			if rowCoincidences > 0 {
				y1 += repeatRange

			} else {
				y1++
			}
		}
		x1++
	}

	return currentCoincidences
}

// GetVerticalCoincidences return nitrogen base vertical coincidences in the given dna matrix
// If macCoincidences > 0 will stop after find maxCoincidences.
func GetVerticalCoincidences(dna []string, repeatRange, maxCoincidences int) int {

	y1 := 0

	maxCoincidencesPerGroup := math.Floor(float64((len(dna[0])) / repeatRange))
	currentCoincidences := 0

	if maxCoincidencesPerGroup == 0 {
		return 0
	}

	fixedRepeatRange := repeatRange - 1

	for y1 < len(dna[0]) {
		rowCoincidences := 0
		x1 := 0
		y2 := y1
		x2 := x1 + int(repeatRange)
		for x1 <= (len(dna[y1]) - repeatRange) {
			x2 = x1 + fixedRepeatRange
			nitrogenBaseCoincidences := 0
			for x2 > x1 {
				xy1 := string(dna[x1][y1])
				xy2 := string(dna[x2][y2])
				if xy1 == xy2 {
					x2--
					nitrogenBaseCoincidences++

				} else {
					break
				}
			}
			if nitrogenBaseCoincidences == fixedRepeatRange {
				rowCoincidences++
				currentCoincidences++
				if currentCoincidences == maxCoincidences {
					return currentCoincidences
				}

				if rowCoincidences == int(maxCoincidencesPerGroup) {
					break
				}

			}
			if rowCoincidences > 0 {
				x1 += repeatRange
			} else {
				x1++
			}
		}
		y1++
	}

	return currentCoincidences
}

// GetObliqueUpLeftCoincidences return nitrogen base up left coincidences in the given dna matrix
// If macCoincidences > 0 will stop after find maxCoincidences.
func GetObliqueUpLeftCoincidences(dna []string, repeatRange, maxCoincidences int) int {

	yref := 0

	maxCoincidencesPerGroup := math.Floor(float64((len(dna[0])) / repeatRange))
	currentCoincidences := 0

	if maxCoincidencesPerGroup == 0 {
		return 0
	}

	fixedRepeatRange := repeatRange - 1

	for yref < len(dna[0]) {
		rowCoincidences := 0
		y1 := yref
		x1 := 0
		x2 := y1 + int(repeatRange)
		y2 := y1 + int(repeatRange)
		for y1 <= (len(dna[x1]) - repeatRange) {
			y2 = y1 + fixedRepeatRange
			x2 = x1 + fixedRepeatRange
			nitrogenBaseCoincidences := 0
			for y2 > y1 && x2 > x1 {
				xy1 := string(dna[x1][y1])
				xy2 := string(dna[x2][y2])
				if xy1 == xy2 {
					y2--
					x2--
					nitrogenBaseCoincidences++

				} else {
					break
				}
			}
			if nitrogenBaseCoincidences == fixedRepeatRange {
				rowCoincidences++
				currentCoincidences++
				if currentCoincidences == maxCoincidences {
					return currentCoincidences
				}

				if rowCoincidences == int(maxCoincidencesPerGroup) {
					break
				}

			}
			if rowCoincidences > 0 {
				x1 += repeatRange
				y1 += repeatRange
			} else {
				x1++
				y1++
			}
		}
		yref++
	}

	return currentCoincidences
}
