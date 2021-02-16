package dnausecases

import (
	"testing"

	"github.com/miguelramirez93/mutant_detector/domain/dna/mocks"
	apperrors "github.com/miguelramirez93/mutant_detector/shared/app_errors"
)

var isMutantUcaseInstance = NewIsMutantUseCase()

func TestIsMutantPositiveValue(t *testing.T) {
	isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaCorrectFormatInput)
	if err != nil {
		t.Errorf("spected err to be nil, got %v", err)
	} else if isMutant {
		t.Logf("spected isMutant to be true")
	} else {
		t.Errorf("spected isMutant to be true, got %v", isMutant)
	}
}

func TestIsMutantWrongDimInput(t *testing.T) {
	_, err := isMutantUcaseInstance.Execute(mocks.DnaBadDimInput)
	if err != nil && err.Err == apperrors.ErrBadParamInput {
		t.Logf("spected err to be %v, got %v", apperrors.ErrBadParamInput, err)
	} else {
		t.Errorf("spected err to be %v, got %v", apperrors.ErrBadParamInput, err)
	}

}

func TestIsMutantWrongLeterInput(t *testing.T) {
	_, err := isMutantUcaseInstance.Execute(mocks.DnaWrongLeterInput)
	if err != nil && err.Err == apperrors.ErrBadParamInput {
		t.Logf("spected err to be %v, got %v", apperrors.ErrBadParamInput, err)
	} else {
		t.Errorf("spected err to be %v, got %v", apperrors.ErrBadParamInput, err)
	}

}
