package nitrogenbaseutils_test

import (
	"testing"

	"mutant_detector/dna/mocks"

	nitrogenbaseutils "mutant_detector/dna/utils/nitrogen_base"

	"github.com/stretchr/testify/assert"
)

func TestNitrogenBaseUtils(t *testing.T) {
	t.Run("Should get Horizontal coincidences (no more than 1 per row)", func(t *testing.T) {
		repeatRange := 4
		maxCoincidences := 2
		testCoincidences := nitrogenbaseutils.GetHorizontalCoincidences(mocks.DnaMutantHorizontal, repeatRange, maxCoincidences)
		assert.Equal(t, maxCoincidences, testCoincidences)
	})

	t.Run("Should get Horizontal coincidences (more than 1 per row)", func(t *testing.T) {
		repeatRange := 3
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetHorizontalCoincidences(mocks.DnaMutantHorizontalCase2, repeatRange, maxCoincidences)
		assert.Equal(t, 6, testCoincidences)
	})

	t.Run("Should get Vertical coincidences (no more than 1 per row)", func(t *testing.T) {
		repeatRange := 4
		maxCoincidences := 2
		testCoincidences := nitrogenbaseutils.GetVerticalCoincidences(mocks.DnaMutantVertical, repeatRange, maxCoincidences)
		assert.Equal(t, maxCoincidences, testCoincidences)
	})

	t.Run("Should get Vertical coincidences (more than 1 per row)", func(t *testing.T) {
		repeatRange := 3
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetVerticalCoincidences(mocks.DnaMutantVertical, repeatRange, maxCoincidences)
		assert.Equal(t, 7, testCoincidences)
	})

	t.Run("Should get ObliqueUpRight coincidences (no more than 1 per row)", func(t *testing.T) {
		repeatRange := 4
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetObliqueUpRightCoincidences(mocks.DnaCorrectFormatInput, repeatRange, maxCoincidences, true)
		assert.Equal(t, 1, testCoincidences)
	})

	t.Run("Should get ObliqueUpRight coincidences (no more than 1 per row) case 2", func(t *testing.T) {
		repeatRange := 4
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetObliqueUpRightCoincidences(mocks.DnaObliqueCase2, repeatRange, maxCoincidences, true)
		assert.Equal(t, 2, testCoincidences)
	})

	t.Run("Should get ObliqueUpRight coincidences (no more than 1 per row)", func(t *testing.T) {
		repeatRange := 3
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetObliqueUpRightCoincidences(mocks.DnaObliqueCase3, repeatRange, maxCoincidences, true)
		assert.Equal(t, 3, testCoincidences)
	})

	t.Run("Should get ObliqueUpRight coincidences (without main diagonal)", func(t *testing.T) {
		repeatRange := 3
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetObliqueUpRightCoincidences(mocks.DnaObliqueCase3, repeatRange, maxCoincidences, false)
		assert.Equal(t, 1, testCoincidences)
	})

	t.Run("Should get ObliqueDownLeft coincidences (no more than 1 per row)", func(t *testing.T) {
		repeatRange := 4
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetObliqueDowmLeftCoincidences(mocks.DnaObliqueDownLeft, repeatRange, maxCoincidences, true)
		assert.Equal(t, 1, testCoincidences)
	})

	t.Run("Should get ObliqueDownLeft coincidences (more than 1 per row)", func(t *testing.T) {
		repeatRange := 3
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetObliqueDowmLeftCoincidences(mocks.DnaObliqueDownLeftCase2, repeatRange, maxCoincidences, true)
		assert.Equal(t, 4, testCoincidences)
	})

	t.Run("Should get ObliqueDownLeft coincidences (without main diagonal)", func(t *testing.T) {
		repeatRange := 3
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetObliqueDowmLeftCoincidences(mocks.DnaObliqueDownLeftCase2, repeatRange, maxCoincidences, false)
		assert.Equal(t, 2, testCoincidences)
	})

	t.Run("Should get ObliqueUpLeft coincidences (no more than 1 per row)", func(t *testing.T) {
		repeatRange := 4
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetObliqueUpLeftCoincidences(mocks.DnaObliqueUpLeftCase1, repeatRange, maxCoincidences, true)
		assert.Equal(t, 2, testCoincidences)
	})

	t.Run("Should get ObliqueUpLeft coincidences (more than 1 per row)", func(t *testing.T) {
		repeatRange := 3
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetObliqueUpLeftCoincidences(mocks.DnaObliqueUpLeftCase1, repeatRange, maxCoincidences, true)
		assert.Equal(t, 4, testCoincidences)
	})
	t.Run("Should get ObliqueUpLeft coincidences (without main diagonal)", func(t *testing.T) {
		repeatRange := 3
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetObliqueUpLeftCoincidences(mocks.DnaObliqueUpLeftCase1, repeatRange, maxCoincidences, false)
		assert.Equal(t, 2, testCoincidences)
	})
	t.Run("Should get ObliqueDownRight coincidences (no more than 1 per row)", func(t *testing.T) {
		repeatRange := 4
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetObliqueDownRightCoincidences(mocks.DnaObliqueDownRight, repeatRange, maxCoincidences, true)
		assert.Equal(t, 2, testCoincidences)
	})
	t.Run("Should get ObliqueDownRight coincidences (more than 1 per row)", func(t *testing.T) {
		repeatRange := 3
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetObliqueDownRightCoincidences(mocks.DnaObliqueDownRight, repeatRange, maxCoincidences, true)
		assert.Equal(t, 4, testCoincidences)
	})
	t.Run("Should get ObliqueDownRight coincidences (without main diagonal)", func(t *testing.T) {
		repeatRange := 3
		maxCoincidences := -1
		testCoincidences := nitrogenbaseutils.GetObliqueDownRightCoincidences(mocks.DnaObliqueDownRight, repeatRange, maxCoincidences, false)
		assert.Equal(t, 2, testCoincidences)
	})
}
