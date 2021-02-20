package repository

import (
	apperrors "mutant_detector/domain/app_errors"
	"mutant_detector/domain/dna"

	"gorm.io/gorm"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type gormIsMutantRepository struct {
	db *gorm.DB
}

type dnaReportPostgres struct {
	dna.DnaReport
	DnaMatrix postgres.Jsonb `gorm:"type:jsonb;unique;index"`
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
