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
		t.Logf("spected result equal to 6, got %v", testCoincidences)
	} else {
		t.Errorf("spected result equal to 6, got %v", testCoincidences)
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
		t.Logf("spected result equal to 7, got %v", testCoincidences)
	} else {
		t.Errorf("spected result equal to 7, got %v", testCoincidences)
	}
}

func TestGetObliqueUpLeftCoincidences(t *testing.T) {
	repeatRange := 4
	maxCoincidences := -1
	testCoincidences := GetObliqueUpLeftCoincidences(mocks.DnaCorrectFormatInput, repeatRange, maxCoincidences, true)

	if testCoincidences == 1 {
		t.Logf("spected result equal to 1, got %v", testCoincidences)
	} else {
		t.Errorf("spected result equal to 1, got %v", testCoincidences)
	}
}

func TestGetObliqueUpLeftCoincidencesCase2(t *testing.T) {
	repeatRange := 4
	maxCoincidences := -1
	testCoincidences := GetObliqueUpLeftCoincidences(mocks.DnaObliqueCase2, repeatRange, maxCoincidences, true)

	if testCoincidences == 2 {
		t.Logf("spected result equal to 2, got %v", testCoincidences)
	} else {
		t.Errorf("spected result equal to 2, got %v", testCoincidences)
	}
}

func TestGetObliqueUpLeftCoincidencesCase3(t *testing.T) {
	repeatRange := 3
	maxCoincidences := -1
	testCoincidences := GetObliqueUpLeftCoincidences(mocks.DnaObliqueCase3, repeatRange, maxCoincidences, true)

	if testCoincidences == 3 {
		t.Logf("spected result equal to 3, got %v", testCoincidences)
	} else {
		t.Errorf("spected result equal to 3, got %v", testCoincidences)
	}
}

func TestGetObliqueUpLeftCoincidencesWithOutMainDiag(t *testing.T) {
	repeatRange := 3
	maxCoincidences := -1
	testCoincidences := GetObliqueUpLeftCoincidences(mocks.DnaObliqueCase3, repeatRange, maxCoincidences, false)

	if testCoincidences == 1 {
		t.Logf("spected result equal to 1, got %v", testCoincidences)
	} else {
		t.Errorf("spected result equal to 1, got %v", testCoincidences)
	}
}

func TestGetObliqueDownRightCoincidencesCase1(t *testing.T) {
	repeatRange := 4
	maxCoincidences := -1
	testCoincidences := GetObliqueDowmRightCoincidences(mocks.DnaObliqueDownRight, repeatRange, maxCoincidences, true)

	if testCoincidences == 1 {
		t.Logf("spected result equal to 1, got %v", testCoincidences)
	} else {
		t.Errorf("spected result equal to 1, got %v", testCoincidences)
	}
}

func TestGetObliqueDownRightCoincidencesCase2(t *testing.T) {
	repeatRange := 3
	maxCoincidences := -1
	testCoincidences := GetObliqueDowmRightCoincidences(mocks.DnaObliqueDownRightCase2, repeatRange, maxCoincidences, true)

	if testCoincidences == 4 {
		t.Logf("spected result equal to 4, got %v", testCoincidences)
	} else {
		t.Errorf("spected result equal to 4, got %v", testCoincidences)
	}
}

func TestGetObliqueDownRightCoincidencesWithOutMainDiagonal(t *testing.T) {
	repeatRange := 3
	maxCoincidences := -1
	testCoincidences := GetObliqueDowmRightCoincidences(mocks.DnaObliqueDownRightCase2, repeatRange, maxCoincidences, false)

	if testCoincidences == 2 {
		t.Logf("spected result equal to 2, got %v", testCoincidences)
	} else {
		t.Errorf("spected result equal to 2, got %v", testCoincidences)
	}
}
