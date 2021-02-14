package dnavalidators

import (
	"fmt"

	apperrors "github.com/miguelramirez93/mutants_detector/shared/app_errors"
	arrayutils "github.com/miguelramirez93/mutants_detector/utils/array_utils"
)

var acceptedNotrogenBases = []string{"A", "T", "C", "G"}

// IsValidDna validate if input is a valid dna value
func IsValidDna(dna []string) *apperrors.AppError {
	for _, dnaRow := range dna {
		if len(dnaRow) != len(dna) {
			return &apperrors.AppError{
				Err:         apperrors.ErrBadParamInput,
				Description: "NxM Dimention is not allowed",
			}
		}
		if !arrayutils.IsStringInRange(dnaRow, acceptedNotrogenBases) {
			return &apperrors.AppError{
				Err:         apperrors.ErrBadParamInput,
				Description: fmt.Sprintf("dna data contains nitrogen base that is not in %v", acceptedNotrogenBases),
			}
		}
	}

	return nil
}
