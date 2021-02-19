package domain

import "fmt"

// constants
var (
	AcceptedNitrogenBases = []string{"A", "T", "C", "G"}
	WrongDimInputError    = &AppError{
		Err:         ErrBadParamInput,
		Description: "NxM Dimention is not allowed",
	}
	WrongNitrogenBaseError = &AppError{
		Err:         ErrBadParamInput,
		Description: fmt.Sprintf("dna data contains nitrogen base that is not in %v", AcceptedNitrogenBases),
	}
)

// IsMutantUsecase contract for ismutant ucase
type IsMutantUsecase interface {
	Execute(dna []string) (bool, *AppError)
}
