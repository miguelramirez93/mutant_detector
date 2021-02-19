package models

// IsMutantReqBody struct used for unmarshall isMutant data
type IsMutantReqBody struct {
	Dna []string `json:"dna" binding:"required"`
}
