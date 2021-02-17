package domain

// IsMutantUsecase contract for is mutant ucase
type IsMutantUsecase interface {
	Execute(dna []string) (bool, *AppError)
}
