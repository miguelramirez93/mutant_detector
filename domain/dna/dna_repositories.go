package dna

import apperrors "mutant_detector/domain/app_errors"

// DNAReportRepository contract for DNA report data I/O operations
type DNAReportRepository interface {
	StoreDNAReport(dnaReport DnaReport) (*DnaReport, *apperrors.AppError)
	GetDNAReportsByIsMutantResult(result bool) ([]DnaReport, *apperrors.AppError)
}
