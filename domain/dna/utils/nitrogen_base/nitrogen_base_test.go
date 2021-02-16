package nitrogenbaseutils

import (
	"testing"

	"github.com/miguelramirez93/mutant_detector/domain/dna/mocks"
)

func TestGetHorizontalCoincidences(t *testing.T) {
	repeatRange := 4
	maxCoincidences := 2
	testCoincidences := GetHorizontalCoincidences(mocks.DnaMutantHorizontal, repeatRange, maxCoincidences)

	if testCoincidences == maxCoincidences {
		t.Logf("spected result equal to %v, got %v", 6, testCoincidences)
	} else {
		t.Errorf("spected result equal to %v, got %v", 6, testCoincidences)
	}
}

func TestGetHorizontalCoincidencesCase2(t *testing.T) {
	repeatRange := 3
	maxCoincidences := -1
	testCoincidences := GetHorizontalCoincidences(mocks.DnaMutantHorizontalCase2, repeatRange, maxCoincidences)

	if testCoincidences == 6 {
		t.Logf("spected result equal to %v, got %v", maxCoincidences, testCoincidences)
	} else {
		t.Errorf("spected result equal to %v, got %v", maxCoincidences, testCoincidences)
	}
}

func TestGetVerticalCoincidences(t *testing.T) {
	repeatRange := 4
	maxCoincidences := 2
	testCoincidences := GetVerticalCoincidences(mocks.DnaMutantVertical, repeatRange, maxCoincidences)

	if testCoincidences == maxCoincidences {
		t.Logf("spected result equal to %v, got %v", maxCoincidences, testCoincidences)
	} else {
		t.Errorf("spected result equal to %v, got %v", maxCoincidences, testCoincidences)
	}
}

func TestGetVerticalCoincidencesCase2(t *testing.T) {
	repeatRange := 3
	maxCoincidences := -1
	testCoincidences := GetVerticalCoincidences(mocks.DnaMutantVertical, repeatRange, maxCoincidences)

	if testCoincidences == 7 {
		t.Logf("spected result equal to %v, got %v", maxCoincidences, testCoincidences)
	} else {
		t.Errorf("spected result equal to %v, got %v", maxCoincidences, testCoincidences)
	}
}
