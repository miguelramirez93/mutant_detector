package dnavalidators

import (
	"testing"

	"github.com/miguelramirez93/mutant_detector/domain/dna/mocks"
	apperrors "github.com/miguelramirez93/mutant_detector/shared/app_errors"
)

func TestIsValidDnaSuccessCase(t *testing.T) {
	err := IsValidDna(mocks.DnaCorrectFormatInput)

	if err != nil {
		t.Errorf("spected err to be nil, got %v", err)
	} else {
		t.Logf("spected err to be nil, got %v", err)
	}
}

func TestIsValidDnaWrongDimCase(t *testing.T) {
	err := IsValidDna(mocks.DnaBadDimInput)

	if err == nil {
		t.Errorf("spected err to be %v, got %v", apperrors.ErrBadParamInput, err)
	} else if err.Err == apperrors.ErrBadParamInput {
		t.Logf("spected err to be %v, got %v", apperrors.ErrBadParamInput, err)
	} else {
		t.Errorf("spected err to be %v, got %v", apperrors.ErrBadParamInput, err)
	}
}

func TestIsValidDnaWrongLeterCase(t *testing.T) {
	err := IsValidDna(mocks.DnaWrongLeterInput)

	if err == nil {
		t.Errorf("spected err to be %v, got %v", apperrors.ErrBadParamInput, err)
	} else if err.Err == apperrors.ErrBadParamInput {
		t.Logf("spected err to be %v, got %v", apperrors.ErrBadParamInput, err)
	} else {
		t.Errorf("spected err to be %v, got %v", apperrors.ErrBadParamInput, err)
	}
}
