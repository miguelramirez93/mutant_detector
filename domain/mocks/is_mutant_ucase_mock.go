package mocks

import (
	apperrors "mutant_detector/domain/app_errors"

	"github.com/stretchr/testify/mock"
)

// IsMutantUseCaseMock mocked functions for IsMutantUseCase
type IsMutantUseCaseMock struct {
	mock.Mock
}

// Execute mocked function of Execute ucase method
func (m *IsMutantUseCaseMock) Execute(dna []string) (bool, *apperrors.AppError) {
	args := m.Called(dna)
	var mockedErr *apperrors.AppError
	if args.Get(1) != nil {
		mockedErr, _ = args.Get(1).(*apperrors.AppError)
	}

	return args.Bool(0), mockedErr
}
