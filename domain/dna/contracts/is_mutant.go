package dnacontracts

import apperrors "github.com/miguelramirez93/mutants_detector/shared/app_errors"

// IsMutantUsecase use for delivery communication.
type IsMutantUsecase interface {
	Execute(dna []string) (bool, *apperrors.AppError)
}