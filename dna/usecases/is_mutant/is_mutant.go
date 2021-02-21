package usecases

import (
	"log"
	nitrogenbaseutils "mutant_detector/dna/utils/nitrogen_base"
	"mutant_detector/dna/validators"
	apperrors "mutant_detector/domain/app_errors"
	"mutant_detector/domain/dna"
	dnadomain "mutant_detector/domain/dna"
	queueutils "mutant_detector/shared/utils/queue_utils"
)

type isMutantUsecase struct {
	dnaReportRepository dna.DNAReportRepository
}

// NewIsMutantUseCase creates new intance from Dnacontracts
func NewIsMutantUseCase(dnaReportRepo dna.DNAReportRepository) dna.IsMutantUsecase {
	return &isMutantUsecase{
		dnaReportRepository: dnaReportRepo,
	}
}

func (uc *isMutantUsecase) Execute(dna []string) (bool, *apperrors.AppError) {
	err := validators.IsValidDna(dna)
	if err != nil {
		return false, err
	}

	repeatRange := 4
	minCoincidences := 2

	currentCoincidences := 0

	currentCoincidences += nitrogenbaseutils.GetHorizontalCoincidences(dna, repeatRange, minCoincidences)

	isMutant := false

	defer func() {
		dnaReportToStore := dnadomain.DnaReport{
			Dna:      dna,
			IsMutant: isMutant,
		}

		queueutils.QueueJob(func() {
			if _, err = uc.dnaReportRepository.StoreDNAReport(dnaReportToStore); err != nil {
				log.Println("Error storing DNA data", err)
			}
		})

	}()

	if currentCoincidences >= minCoincidences {
		isMutant = true
		return isMutant, nil
	}

	currentCoincidences += nitrogenbaseutils.GetVerticalCoincidences(dna, repeatRange, minCoincidences)

	if currentCoincidences >= minCoincidences {
		isMutant = true
		return isMutant, nil
	}

	currentCoincidences += nitrogenbaseutils.GetObliqueUpLeftCoincidences(dna, repeatRange, minCoincidences, true)

	if currentCoincidences >= minCoincidences {
		isMutant = true
		return isMutant, nil
	}

	currentCoincidences += nitrogenbaseutils.GetObliqueDownRightCoincidences(dna, repeatRange, minCoincidences, false)

	if currentCoincidences >= minCoincidences {
		isMutant = true
		return isMutant, nil
	}

	currentCoincidences += nitrogenbaseutils.GetObliqueUpRightCoincidences(dna, repeatRange, minCoincidences, true)

	if currentCoincidences >= minCoincidences {
		isMutant = true
		return isMutant, nil
	}

	currentCoincidences += nitrogenbaseutils.GetObliqueDowmLeftCoincidences(dna, repeatRange, minCoincidences, false)

	if currentCoincidences >= minCoincidences {
		isMutant = true
		return isMutant, nil
	}

	return isMutant, nil
}
