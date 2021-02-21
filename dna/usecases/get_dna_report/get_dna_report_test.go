package usecases_test

import (
	getdnareportucase "mutant_detector/dna/usecases/get_dna_report"
	"mutant_detector/domain/dna"
	domainmocks "mutant_detector/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDNAReport(t *testing.T) {
	t.Run("Should return 0", func(t *testing.T) {
		dnaReportSpected := []dna.DnaReport{}
		dnaReportRepoMock := new(domainmocks.DNAReportRepositoryMock)
		dnaReportRepoMock.On("GetDNAReportsByIsMutantResult", true).Once().Return(dnaReportSpected, nil)
		dnaReportRepoMock.On("GetDNAReportsByIsMutantResult", false).Once().Return(dnaReportSpected, nil)

		getDnaReportInstance := getdnareportucase.NewGetDnaReportUseCase(dnaReportRepoMock)

		statResult, err := getDnaReportInstance.Execute()
		assert.Equal(t, &dna.DnaReportStats{
			CountMutantDNA: 0,
			CountHumanDNA:  0,
			Ratio:          0,
		}, statResult)

		assert.Nil(t, err)
	})

	t.Run("Should return 0", func(t *testing.T) {
		dnaReportSpected := []dna.DnaReport{
			dna.DnaReport{
				ID:       1,
				Dna:      []string{},
				IsMutant: true,
			},
		}
		dnaReportRepoMock := new(domainmocks.DNAReportRepositoryMock)
		dnaReportRepoMock.On("GetDNAReportsByIsMutantResult", true).Once().Return(dnaReportSpected, nil)
		dnaReportRepoMock.On("GetDNAReportsByIsMutantResult", false).Once().Return(dnaReportSpected, nil)

		getDnaReportInstance := getdnareportucase.NewGetDnaReportUseCase(dnaReportRepoMock)

		statResult, err := getDnaReportInstance.Execute()
		assert.Equal(t, &dna.DnaReportStats{
			CountMutantDNA: 1,
			CountHumanDNA:  1,
			Ratio:          1,
		}, statResult)

		assert.Nil(t, err)
	})
}
