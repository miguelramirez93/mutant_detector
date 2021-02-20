package dna

import (
	"database/sql/driver"
	"encoding/json"
)

// DnaType type for dna input (used by gorm)
type DnaType []string

// Value used by gorm
func (j DnaType) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

// Scan used by gorm
func (j *DnaType) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

// AcceptedNitrogenBases current accepted NB for dna matrix
var AcceptedNitrogenBases = []string{"A", "T", "C", "G"}

// DnaReport struct that represents data that will be stored with repository
type DnaReport struct {
	ID       int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Dna      DnaType `json:"dna" gorm:"unique;index;type:jsonb"`
	IsMutant bool    `json:"isMutant" gorm:"column:is_mutant"`
}

// DnaReportSTats represents report stats for dna.
type DnaReportStats struct {
	CountMutantDNA int     `json:"count_mutant_dna"`
	CountHumanDNA  int     `json:"count_human_dna"`
	Ratio          float64 `json:"ratio"`
}
