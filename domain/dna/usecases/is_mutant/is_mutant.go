package dnausecases

import (
	dnacontracts "github.com/miguelramirez93/mutant_detector/domain/dna/contracts"
	nitrogenbaseutils "github.com/miguelramirez93/mutant_detector/domain/dna/utils/nitrogen_base"
	dnavalidators "github.com/miguelramirez93/mutant_detector/domain/dna/validators"
	apperrors "github.com/miguelramirez93/mutant_detector/shared/app_errors"
)

type isMutantUsecase struct{}

// NewIsMutantUseCase creates new intance from Dnacontracts
func NewIsMutantUseCase() dnacontracts.IsMutantUsecase {
	return &isMutantUsecase{}
}

func (uc *isMutantUsecase) Execute(dna []string) (bool, *apperrors.AppError) {
	err := dnavalidators.IsValidDna(dna)
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
