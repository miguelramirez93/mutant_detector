package mocks

var (
	//DnaCorrectFormatInput dna data with spected format
	DnaCorrectFormatInput = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	// DnaBadDimInput dna with nxm format
	DnaBadDimInput = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACT"}
	// DnaWrongLeterInput dna with nitrogen base out of the main domain
	DnaWrongLeterInput = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACH"}
)
