package usecases_test

import (
	"fmt"
	"testing"

	"mutant_detector/dna/mocks"
	"mutant_detector/domain"

	usecases "mutant_detector/dna/usecases/is_mutant"

	"github.com/stretchr/testify/assert"
)

var isMutantUcaseInstance = usecases.NewIsMutantUseCase()

func TestIsMutantUCase(t *testing.T) {
	t.Run("Should return true for a expected mutant dna input", func(t *testing.T) {
		isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaCorrectFormatInput)
		fmt.Println(err)
		assert.Nil(t, err)
		assert.Equal(t, true, isMutant)
	})
	t.Run("Should return false for a expected human dna input", func(t *testing.T) {
		isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaHumanInput)
		fmt.Println(err)
		assert.Nil(t, err)
		assert.Equal(t, false, isMutant)
	})
	t.Run("Should return error when dna matrix is not nxn", func(t *testing.T) {
		isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaBadDimInput)
		assert.Equal(t, domain.WrongDimInputError, err)
		assert.Equal(t, false, isMutant)
	})
	t.Run("Should return error when dna matrix contain wrong nitrogen bases", func(t *testing.T) {
		isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaWrongLeterInput)
		assert.Equal(t, domain.WrongNitrogenBaseError, err)
		assert.Equal(t, false, isMutant)
	})
	t.Run("Should return true with horizontal only coincidences matrix", func(t *testing.T) {
		isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaMutantHorizontal)
		assert.Nil(t, err)
		assert.Equal(t, true, isMutant)
	})

	t.Run("Should return true with ObliqueUpLeft only coincidences matrix", func(t *testing.T) {
		isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaObliqueUpLeftCase1)
		assert.Nil(t, err)
		assert.Equal(t, true, isMutant)
	})

	t.Run("Should return true with ObliqueDownRight only coincidences matrix", func(t *testing.T) {
		isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaObliqueDownRight)
		assert.Nil(t, err)
		assert.Equal(t, true, isMutant)
	})

	t.Run("Should return true with ObliqueUpRight only coincidences matrix", func(t *testing.T) {
		isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaObliqueCase2)
		assert.Nil(t, err)
		assert.Equal(t, true, isMutant)
	})

	t.Run("Should return true with ObliqueDownLeft only coincidences matrix", func(t *testing.T) {
		isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaObliqueDownLeft)
		assert.Nil(t, err)
		assert.Equal(t, true, isMutant)
	})
}
