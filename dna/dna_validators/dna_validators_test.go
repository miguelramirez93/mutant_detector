package dnavalidators

import (
	"testing"

	"github.com/miguelramirez93/mutant_detector/dna/mocks"
	domain "github.com/miguelramirez93/mutant_detector/domain"
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
		t.Errorf("spected err to be %v, got %v", domain.ErrBadParamInput, err)
	} else if err.Err == domain.ErrBadParamInput {
		t.Logf("spected err to be %v, got %v", domain.ErrBadParamInput, err)
	} else {
		t.Errorf("spected err to be %v, got %v", domain.ErrBadParamInput, err)
	}
}

func TestIsValidDnaWrongLeterCase(t *testing.T) {
	err := IsValidDna(mocks.DnaWrongLeterInput)

	if err == nil {
		t.Errorf("spected err to be %v, got %v", domain.ErrBadParamInput, err)
	} else if err.Err == domain.ErrBadParamInput {
		t.Logf("spected err to be %v, got %v", domain.ErrBadParamInput, err)
	} else {
		t.Errorf("spected err to be %v, got %v", domain.ErrBadParamInput, err)
	}
}
