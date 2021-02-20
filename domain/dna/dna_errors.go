package dna

import (
	"fmt"
	apperrors "mutant_detector/domain/app_errors"
)

var (

	// WrongDimInputError dnaerror
	WrongDimInputError = &apperrors.AppError{
		Err:         apperrors.ErrBadParamInput,
		Description: "NxM Dimention is not allowed",
	}
	// WrongNitrogenBaseError dnaerror
	WrongNitrogenBaseError = &apperrors.AppError{
		Err:         apperrors.ErrBadParamInput,
		Description: fmt.Sprintf("dna data contains nitrogen base that is not in %v", AcceptedNitrogenBases),
	}
)
