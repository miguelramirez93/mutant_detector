package http

import (
	"mutants_detector/dna/delivery/http/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miguelramirez93/mutant_detector/domain"
)

// DnaHandler http handler struct for dna use cases
type DnaHandler struct {
	IsMutantUcase domain.IsMutantUsecase
}

// NewDnaHandler creates a new instance of DnaHandler
func NewDnaHandler(r *gin.Engine, isMutantUcase domain.IsMutantUsecase) {
	handler := &DnaHandler{
		IsMutantUcase: isMutantUcase,
	}

	r.POST("/mutant", handler.IsMutant)
}

// IsMutant handlr for http request to isMutant use case
func (h *DnaHandler) IsMutant(c *gin.Context) {
	var (
		req      models.IsMutantReqBody
		appError *domain.AppError
		isMutant bool
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		appError = &domain.AppError{
			Err:         domain.ErrBadParamInput,
			Description: err.Error(),
		}
		c.JSON(http.StatusBadRequest, appError)
		return
	}

	isMutant, appError = h.IsMutantUcase.Execute(req.Dna)

	if appError != nil {
		c.JSON(http.StatusBadRequest, appError)
		return
	}

	if !isMutant {
		c.JSON(http.StatusForbidden, nil)
	}

	c.JSON(http.StatusOK, nil)
	return
}
