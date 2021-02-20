package mocks

import (
	apperrors "mutant_detector/domain/app_errors"
	"mutant_detector/domain/dna"

	"github.com/stretchr/testify/mock"
)

// GetDNAReportsStatsUseCaseMock mocked functions for GetDNAReportsStatsUseCase
type GetDNAReportsStatsUseCaseMock struct {
	mock.Mock
}

// Execute mocked function of Execute ucase method
func (m *GetDNAReportsStatsUseCaseMock) Execute() (*dna.DnaReportStats, *apperrors.AppError) {
	args := m.Called()
	var mockedErr *apperrors.AppError
	var mockedResult *dna.DnaReportStats
	if args.Get(0) != nil {
		mockedResult, _ = args.Get(0).(*dna.DnaReportStats)
	}
	if args.Get(1) != nil {
		mockedErr, _ = args.Get(1).(*apperrors.AppError)
	}

	return mockedResult, mockedErr
}
