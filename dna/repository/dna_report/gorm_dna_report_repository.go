package repository

import (
	apperrors "mutant_detector/domain/app_errors"
	"mutant_detector/domain/dna"

	"gorm.io/gorm"
)

type gormIsMutantRepository struct {
	db *gorm.DB
}

// NewGormIsMutantRepository creates a new instance of IsMutantRepo for gorm provider
func NewGormIsMutantRepository(db interface{}) dna.DNAReportRepository {
	return &gormIsMutantRepository{
		db: db.(*gorm.DB),
	}
}

// StoreDNAReport
func (r *gormIsMutantRepository) StoreDNAReport(dnaReport dna.DnaReport) (*dna.DnaReport, *apperrors.AppError) {
	result := r.db.Create(&dnaReport)

	if result.Error != nil {
		return nil, &apperrors.AppError{
			Err:         apperrors.ErrStoringData,
			Description: result.Error.Error(),
		}
	}

	return &dnaReport, nil
}

func (r *gormIsMutantRepository) GetDNAReportsByIsMutantResult(resultToSearch bool) ([]dna.DnaReport, *apperrors.AppError) {
	var foundReports []dna.DnaReport
	m := make(map[string]interface{})
	m["is_mutant"] = resultToSearch
	result := r.db.Where(m).Find(&foundReports)
	if result.Error != nil {
		return nil, &apperrors.AppError{
			Err:         result.Error,
			Description: "Error getting report data",
		}
	}
	return foundReports, nil
}
