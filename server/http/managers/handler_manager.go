package managers

import (
	"fmt"
	"mutant_detector/dna/delivery/http"
	repository "mutant_detector/dna/repository/dna_report"
	getdnareportstatsucase "mutant_detector/dna/usecases/get_dna_report"
	ismutantucase "mutant_detector/dna/usecases/is_mutant"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// HTTPInitDocumentation init http api documentation
func HTTPInitDocumentation(r *gin.Engine) {
	baseURL, ok := os.LookupEnv("HTTP_DOMAIN")
	if !ok {
		panic("http-docs-error")
	}
	url := ginSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", baseURL))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

// HTTPInitHandlers handlers initialitation with coresponding usecases for http delivery
func HTTPInitHandlers(r *gin.Engine, dbConnInstance interface{}) {
	//Repository instances
	dnaReportRepository := repository.NewGormIsMutantRepository(dbConnInstance)

	// UserCases instanses
	isMutantUseCase := ismutantucase.NewIsMutantUseCase(dnaReportRepository)
	getDNAReportStats := getdnareportstatsucase.NewGetDnaReportUseCase(dnaReportRepository)

	// Init handlers
	http.NewDnaHandler(r, isMutantUseCase, getDNAReportStats)
	HTTPInitDocumentation(r)

}
