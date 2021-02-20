package validators

import (
	apperrors "mutant_detector/domain/app_errors"
	dnadomain "mutant_detector/domain/dna"
	arrayutils "mutant_detector/shared/utils/array_utils"
)

// IsValidDna validate if input is a valid dna value
func IsValidDna(dna []string) *apperrors.AppError {
	for _, dnaRow := range dna {
		if len(dnaRow) != len(dna) {
			return dnadomain.WrongDimInputError
		}
		if !arrayutils.IsStringInRange(dnaRow, dnadomain.AcceptedNitrogenBases) {
			return dnadomain.WrongNitrogenBaseError
		}
	}

	return nil
}
