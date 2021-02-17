package mocks

var (
	//DnaCorrectFormatInput dna data with spected format
	DnaCorrectFormatInput = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	// DnaBadDimInput dna with nxm format
	DnaBadDimInput = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACT"}
	// DnaWrongLeterInput dna with nitrogen base out of the main domain
	DnaWrongLeterInput = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACH"}
	// DnaHumanNoInRange dns with n less than repeat range (number of equal nitrogen bases)
	DnaHumanNoInRange = []string{"ATG", "CAG", "TTA"}
	// DnaMutantHorizontal dns with only horizontal coincidences
	DnaMutantHorizontal = []string{"AAAAGT", "CCCCGC", "TTGGGG", "ATTTTG", "CCCCTA", "TCACTG"}
	// DnaMutantHorizontalCase2 dns with only horizontal coincidences
	DnaMutantHorizontalCase2 = []string{"AAAAAA", "CCCCGC", "TTGGGG", "ATTTTG", "CCCCTA", "TCACTG"}
	// DnaMutantVertical dns with only vertical coincidences
	DnaMutantVertical = []string{"ACCTGT", "ACCCGC", "ACGGGG", "ACTTTG", "ACCCTG", "ACACTG"}
	//DnaObliqueCase2 dna data with spected format
	DnaObliqueCase2 = []string{"ATGCGA", "CATTGC", "TTATGT", "AGAATG", "CCCCTA", "TCACTG"}
	//DnaObliqueCase3 dna data with spected format
	DnaObliqueCase3 = []string{"ATGCGA", "CATTGC", "TTATGT", "AGAATG", "CCCCAA", "TCACTA"}
	//DnaObliqueDownLeft dna data with spected format
	DnaObliqueDownLeft = []string{"AAAAAA", "ACCCGC", "TAGGGG", "ATATTG", "CCCATA", "TCACTG"}
	//DnaObliqueDownLeftCase2 dna data with spected format
	DnaObliqueDownLeftCase2 = []string{"GAAAAA", "AGCCGC", "TAGGGG", "ATAGTG", "CCTAGA", "TCATTG"}
	//DnaObliqueUpLeftCase1 dna data with spected format
	DnaObliqueUpLeftCase1 = []string{"ACGTAG", "CGTAGC", "GTTGAC", "TGGTAA", "CGGTAA", "GCGATA"}
	//DnaObliqueDownRight dna data with spected format
	DnaObliqueDownRight = []string{"ACTGAG", "CTGAGT", "CTGGTA", "CTGTGA", "AGTCAT", "GAAACT"}
)
