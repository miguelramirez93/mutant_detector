package managers

import (
	"mutant_detector/dna/delivery/http"
	usecases "mutant_detector/dna/usecases/is_mutant"

	"github.com/gin-gonic/gin"
)

// HTTPInitHandlers handlers initialitation with coresponding usecases for http delivery
func HTTPInitHandlers(r *gin.Engine) {
	// UserCases instanses
	isMutantUseCase := usecases.NewIsMutantUseCase()

	// Init handlers
	http.NewDnaHandler(r, isMutantUseCase)

}
