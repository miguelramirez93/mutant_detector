package mocks

import (
	"mutant_detector/domain"

	"github.com/stretchr/testify/mock"
)

// IsMutantUseCaseMock mocked functions for IsMutantUseCase
type IsMutantUseCaseMock struct {
	mock.Mock
}

// Execute mocked function of Execute ucase method
func (m *IsMutantUseCaseMock) Execute(dna []string) (bool, *domain.AppError) {
	args := m.Called(dna)
	var mockedErr *domain.AppError
	if args.Get(1) != nil {
		mockedErr, _ = args.Get(1).(*domain.AppError)
	}

	return args.Bool(0), mockedErr
}
