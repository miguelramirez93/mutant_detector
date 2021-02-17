package nitrogenbaseutils

import "math"

// GetHorizontalCoincidences return nitrogen base horizontal coincidences in the given dna matrix
// If maxCoincidences > 0 will stop after find maxCoincidences.
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
		y2 := y1 + repeatRange
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
// If maxCoincidences > 0 will stop after find maxCoincidences.
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
		x2 := x1 + repeatRange
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

// GetObliqueUpRightCoincidences return nitrogen base up right coincidences in the given dna matrix
// If maxCoincidences > 0 will stop after find maxCoincidences.
func GetObliqueUpRightCoincidences(dna []string, repeatRange, maxCoincidences int, withMainDiag bool) int {

	yref := 1

	if withMainDiag {
		yref = 0
	}

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
		x2 := y1 + repeatRange
		y2 := y1 + repeatRange
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

// GetObliqueDowmLeftCoincidences return nitrogen base down left coincidences in the given dna matrix
// If maxCoincidences > 0 will stop after find maxCoincidences.
func GetObliqueDowmLeftCoincidences(dna []string, repeatRange, maxCoincidences int, withMainDiag bool) int {

	xref := 1

	if withMainDiag {
		xref = 0
	}

	maxCoincidencesPerGroup := math.Floor(float64((len(dna[0])) / repeatRange))
	currentCoincidences := 0

	if maxCoincidencesPerGroup == 0 {
		return 0
	}

	fixedRepeatRange := repeatRange - 1

	for xref < len(dna) {
		rowCoincidences := 0
		x1 := xref
		y1 := 0
		x2 := y1 + repeatRange
		y2 := y1 + repeatRange
		for y1 <= (len(dna[0])-repeatRange) && x1 <= (len(dna)-repeatRange) {
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
		xref++
	}

	return currentCoincidences
}

// GetObliqueUpLeftCoincidences return nitrogen base up left coincidences in the given dna matrix
// If maxCoincidences > 0 will stop after find maxCoincidences.
func GetObliqueUpLeftCoincidences(dna []string, repeatRange, maxCoincidences int, withMainDiag bool) int {

	yref := len(dna[0]) - 2

	if withMainDiag {
		yref++
	}

	maxCoincidencesPerGroup := math.Floor(float64((len(dna[0])) / repeatRange))
	currentCoincidences := 0

	if maxCoincidencesPerGroup == 0 {
		return 0
	}

	fixedRepeatRange := repeatRange - 1

	for yref >= 0 {
		rowCoincidences := 0
		y1 := yref
		x1 := 0
		x2 := y1 + repeatRange
		y2 := y1 - repeatRange
		for (y1 - fixedRepeatRange) >= 0 {
			y2 = y1 - fixedRepeatRange
			x2 = x1 + fixedRepeatRange
			nitrogenBaseCoincidences := 0
			for y2 < y1 && x2 > x1 {
				xy1 := string(dna[x1][y1])
				xy2 := string(dna[x2][y2])
				if xy1 == xy2 {
					y2++
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
				y1 -= repeatRange
			} else {
				x1++
				y1--
			}
		}
		yref--
	}

	return currentCoincidences
}

// GetObliqueDownRightCoincidences return nitrogen base down right coincidences in the given dna matrix
// If maxCoincidences > 0 will stop after find maxCoincidences.
func GetObliqueDownRightCoincidences(dna []string, repeatRange, maxCoincidences int, withMainDiag bool) int {

	xref := 1

	if withMainDiag {
		xref = 0
	}

	maxCoincidencesPerGroup := math.Floor(float64((len(dna[0])) / repeatRange))
	currentCoincidences := 0

	if maxCoincidencesPerGroup == 0 {
		return 0
	}

	fixedRepeatRange := repeatRange - 1

	for xref < len(dna) {
		rowCoincidences := 0
		x1 := xref
		y1 := len(dna[0]) - 1
		x2 := y1 + repeatRange
		y2 := y1 - repeatRange
		for x1 <= (len(dna) - repeatRange) {
			y2 = y1 - fixedRepeatRange
			x2 = x1 + fixedRepeatRange
			nitrogenBaseCoincidences := 0
			for y2 < y1 && x2 > x1 {
				xy1 := string(dna[x1][y1])
				xy2 := string(dna[x2][y2])
				if xy1 == xy2 {
					y2++
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
				y1 -= repeatRange
			} else {
				x1++
				y1--
			}
		}
		xref++
	}

	return currentCoincidences
}
