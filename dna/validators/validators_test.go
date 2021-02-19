package validators_test

import (
	"testing"

	"mutant_detector/dna/mocks"
	"mutant_detector/dna/validators"
	domain "mutant_detector/domain"

	"github.com/stretchr/testify/assert"
)

func TestValidators(t *testing.T) {
	t.Run("Should return nil err if is valid dna", func(t *testing.T) {
		err := validators.IsValidDna(mocks.DnaCorrectFormatInput)
		assert.Nil(t, err)
	})

	t.Run("Should return bad param input error if dna matrix is not NXN", func(t *testing.T) {
		err := validators.IsValidDna(mocks.DnaBadDimInput)
		assert.Equal(t, domain.WrongDimInputError, err)
	})

	t.Run("Should return bad param input error if dna matrix is not NXN", func(t *testing.T) {
		err := validators.IsValidDna(mocks.DnaWrongLeterInput)
		assert.Equal(t, domain.WrongNitrogenBaseError, err)
	})
}
