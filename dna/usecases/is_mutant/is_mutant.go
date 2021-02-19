package usecases

import (
	nitrogenbaseutils "mutant_detector/dna/utils/nitrogen_base"
	"mutant_detector/dna/validators"
	"mutant_detector/domain"
)

type isMutantUsecase struct{}

// NewIsMutantUseCase creates new intance from Dnacontracts
func NewIsMutantUseCase() domain.IsMutantUsecase {
	return &isMutantUsecase{}
}

func (uc *isMutantUsecase) Execute(dna []string) (bool, *domain.AppError) {
	err := validators.IsValidDna(dna)
	if err != nil {
		return false, err
	}

	repeatRange := 4
	minCoincidences := 2

	currentCoincidences := 0

	currentCoincidences += nitrogenbaseutils.GetHorizontalCoincidences(dna, repeatRange, minCoincidences)

	if currentCoincidences >= minCoincidences {
		return true, nil
	}

	currentCoincidences += nitrogenbaseutils.GetVerticalCoincidences(dna, repeatRange, minCoincidences)

	if currentCoincidences >= minCoincidences {
		return true, nil
	}

	currentCoincidences += nitrogenbaseutils.GetObliqueUpLeftCoincidences(dna, repeatRange, minCoincidences, true)

	if currentCoincidences >= minCoincidences {
		return true, nil
	}

	currentCoincidences += nitrogenbaseutils.GetObliqueDownRightCoincidences(dna, repeatRange, minCoincidences, false)

	if currentCoincidences >= minCoincidences {
		return true, nil
	}

	currentCoincidences += nitrogenbaseutils.GetObliqueUpRightCoincidences(dna, repeatRange, minCoincidences, true)

	if currentCoincidences >= minCoincidences {
		return true, nil
	}

	currentCoincidences += nitrogenbaseutils.GetObliqueDowmLeftCoincidences(dna, repeatRange, minCoincidences, false)

	if currentCoincidences >= minCoincidences {
		return true, nil
	}

	return false, nil
}
