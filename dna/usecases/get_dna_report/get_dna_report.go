package usecases

import (
	apperrors "mutant_detector/domain/app_errors"
	"mutant_detector/domain/dna"
)

type getDnaReportUsecase struct {
	dnaReportRepository dna.DNAReportRepository
}

// NewGetDnaReportUseCase creates new intance from Dnacontracts
func NewGetDnaReportUseCase(dnaReportRepo dna.DNAReportRepository) dna.GetDNAReportsStatsUseCase {
	return &getDnaReportUsecase{
		dnaReportRepository: dnaReportRepo,
	}
}

func (uc *getDnaReportUsecase) Execute() (*dna.DnaReportStats, *apperrors.AppError) {
	mutantReports, err := uc.dnaReportRepository.GetDNAReportsByIsMutantResult(true)
	if err != nil {
		return nil, err
	}

	humanReports, err := uc.dnaReportRepository.GetDNAReportsByIsMutantResult(false)
	if err != nil {
		return nil, err
	}

	totalHumans := len(humanReports)

	totalMutants := len(mutantReports)

	ratio := 0.0

	if totalMutants > 0 {
		ratio = (float64(totalHumans) / float64(totalMutants))
	}

	statsResult := &dna.DnaReportStats{
		CountMutantDNA: totalMutants,
		CountHumanDNA:  totalHumans,
		Ratio:          ratio,
	}

	return statsResult, nil
}
