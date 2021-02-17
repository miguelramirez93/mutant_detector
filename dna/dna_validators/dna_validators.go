package dnavalidators

import (
	"fmt"

	domain "github.com/miguelramirez93/mutant_detector/domain"
	arrayutils "github.com/miguelramirez93/mutant_detector/shared/utils/array_utils"
)

var acceptedNotrogenBases = []string{"A", "T", "C", "G"}

// IsValidDna validate if input is a valid dna value
func IsValidDna(dna []string) *domain.AppError {
	for _, dnaRow := range dna {
		if len(dnaRow) != len(dna) {
			return &domain.AppError{
				Err:         domain.ErrBadParamInput,
				Description: "NxM Dimention is not allowed",
			}
		}
		if !arrayutils.IsStringInRange(dnaRow, acceptedNotrogenBases) {
			return &domain.AppError{
				Err:         domain.ErrBadParamInput,
				Description: fmt.Sprintf("dna data contains nitrogen base that is not in %v", acceptedNotrogenBases),
			}
		}
	}

	return nil
}
