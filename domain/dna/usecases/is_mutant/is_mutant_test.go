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

func TestIsMutantNegativeValueNLessThanRange(t *testing.T) {
	isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaHumanNoInRange)
	if err != nil {
		t.Errorf("spected err to be nil, got %v", err)
	} else if !isMutant {
		t.Logf("spected isMutant to be false, dna matrix is nxn with n < repeatRange")
	} else {
		t.Errorf("spected isMutant to be false, dna matrix is nxn with n < repeatRange, got %v", isMutant)
	}
}

func TestIsMutantPositiveValueHorizontalOnlyCases(t *testing.T) {
	isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaMutantHorizontal)
	if err != nil {
		t.Errorf("spected err to be nil, got %v", err)
	} else if isMutant {
		t.Logf("spected isMutant to be true with only horizontal coincidences")
	} else {
		t.Errorf("spected isMutant to be true with only horizontal coincidences, got %v", isMutant)
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

func TestIsMutantPositiveValueObliqueUpLeftOnlyCases(t *testing.T) {
	isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaObliqueUpLeftCase1)
	if err != nil {
		t.Errorf("spected err to be nil, got %v", err)
	} else if isMutant {
		t.Logf("spected isMutant to be true with only ObliqueUpLeft coincidences")
	} else {
		t.Errorf("spected isMutant to be true with only ObliqueUpLeft coincidences, got %v", isMutant)
	}
}

func TestIsMutantPositiveValueObliqueDownRightOnlyCases(t *testing.T) {
	isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaObliqueDownRight)
	if err != nil {
		t.Errorf("spected err to be nil, got %v", err)
	} else if isMutant {
		t.Logf("spected isMutant to be true with only DownRight coincidences")
	} else {
		t.Errorf("spected isMutant to be true with only DownRight coincidences, got %v", isMutant)
	}
}

func TestIsMutantPositiveValueObliqueUpRightOnlyCases(t *testing.T) {
	isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaObliqueCase2)
	if err != nil {
		t.Errorf("spected err to be nil, got %v", err)
	} else if isMutant {
		t.Logf("spected isMutant to be true with only UpRight coincidences")
	} else {
		t.Errorf("spected isMutant to be true with only UpRight coincidences, got %v", isMutant)
	}
}

func TestIsMutantPositiveValueObliqueDownLeftOnlyCases(t *testing.T) {
	isMutant, err := isMutantUcaseInstance.Execute(mocks.DnaObliqueDownLeft)
	if err != nil {
		t.Errorf("spected err to be nil, got %v", err)
	} else if isMutant {
		t.Logf("spected isMutant to be true with only DownLeft coincidences")
	} else {
		t.Errorf("spected isMutant to be true with only DownLeft coincidences, got %v", isMutant)
	}
}
