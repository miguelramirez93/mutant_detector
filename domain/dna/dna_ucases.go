package dna

import apperrors "mutant_detector/domain/app_errors"

// IsMutantUsecase contract for ismutant ucase
type IsMutantUsecase interface {
	Execute(dna []string) (bool, *apperrors.AppError)
}

// GetDNAReportsStatsUseCase contract for get reports stats ucase
type GetDNAReportsStatsUseCase interface {
	Execute() (*DnaReportStats, *apperrors.AppError)
}
