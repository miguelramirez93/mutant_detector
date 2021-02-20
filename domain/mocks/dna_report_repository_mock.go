package mocks

import (
	"mutant_detector/dna/mocks"
	apperrors "mutant_detector/domain/app_errors"
	"mutant_detector/domain/dna"

	"github.com/stretchr/testify/mock"
)

// DNAReportRepositoryMock mocked functions for IsMutantRepository
type DNAReportRepositoryMock struct {
	mock.Mock
}

// StoreDNAReport mocked function of StoreDNAReport repository method
func (m *DNAReportRepositoryMock) StoreDNAReport(dnaReport dna.DnaReport) (*dna.DnaReport, *apperrors.AppError) {
	args := m.Called(dnaReport)
	var mockedErr *apperrors.AppError
	var mockedResult *dna.DnaReport
	if args.Get(0) != nil {
		mockedResult, _ = args.Get(0).(*dna.DnaReport)
	}
	if args.Get(1) != nil {
		mockedErr, _ = args.Get(1).(*apperrors.AppError)
	}
	return mockedResult, mockedErr
}

// GetDNAReportsByIsMutantResult mocked function of GetDNAReportsByIsMutantResult repository method
func (m *DNAReportRepositoryMock) GetDNAReportsByIsMutantResult(result bool) ([]dna.DnaReport, *apperrors.AppError) {
	args := m.Called(result)
	var mockedErr *apperrors.AppError
	var mockedResult []dna.DnaReport
	if args.Get(0) != nil {
		mockedResult, _ = args.Get(0).([]dna.DnaReport)
	}
	if args.Get(1) != nil {
		mockedErr, _ = args.Get(1).(*apperrors.AppError)
	}
	return mockedResult, mockedErr
}

var (
	StoreDNAReportSuccess = dna.DnaReport{
		Dna:      mocks.DnaCorrectFormatInput,
		IsMutant: true,
	}
)
