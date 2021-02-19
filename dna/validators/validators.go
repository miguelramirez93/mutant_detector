package validators

import (
	domain "mutant_detector/domain"
	arrayutils "mutant_detector/shared/utils/array_utils"
)

// IsValidDna validate if input is a valid dna value
func IsValidDna(dna []string) *domain.AppError {
	for _, dnaRow := range dna {
		if len(dnaRow) != len(dna) {
			return domain.WrongDimInputError
		}
		if !arrayutils.IsStringInRange(dnaRow, domain.AcceptedNitrogenBases) {
			return domain.WrongNitrogenBaseError
		}
	}

	return nil
}
