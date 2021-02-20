package dna

import apperrors "mutant_detector/domain/app_errors"

// IsMutantUsecase contract for ismutant ucase
type IsMutantUsecase interface {
	Execute(dna []string) (bool, *apperrors.AppError)
}
