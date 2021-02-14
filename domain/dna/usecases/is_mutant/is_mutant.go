package dnausecases

import (
	dnacontracts "github.com/miguelramirez93/mutants_detector/domain/dna/contracts"
	dnavalidators "github.com/miguelramirez93/mutants_detector/domain/dna/validators"
	apperrors "github.com/miguelramirez93/mutants_detector/shared/app_errors"
)

type isMutantUsecase struct{}

// NewIsMutantUseCase creates new intance from DnaCheckUsecase contract
func NewIsMutantUseCase() dnacontracts.IsMutantUsecase {
	return &isMutantUsecase{}
}

func (uc *isMutantUsecase) Execute(dna []string) (bool, *apperrors.AppError) {
	err := dnavalidators.IsValidDna(dna)
	if err != nil {
		return false, err
	}

	return false, nil
}
