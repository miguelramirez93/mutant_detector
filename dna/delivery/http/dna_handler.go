package http

import (
	"mutant_detector/dna/delivery/http/models"
	apperrors "mutant_detector/domain/app_errors"
	"mutant_detector/domain/dna"
	"mutant_detector/shared/mappers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DnaHandler http handler struct for dna use cases
type DnaHandler struct {
	IsMutantUcase          dna.IsMutantUsecase
	GetDnaReportStatsUcase dna.GetDNAReportsStatsUseCase
}

// NewDnaHandler creates a new instance of DnaHandler with corresponding routes
func NewDnaHandler(r *gin.Engine, isMutantUcase dna.IsMutantUsecase, getDnaReportUsecase dna.GetDNAReportsStatsUseCase) {
	handler := &DnaHandler{
		IsMutantUcase:          isMutantUcase,
		GetDnaReportStatsUcase: getDnaReportUsecase,
	}

	r.POST("/mutant", handler.IsMutant)
	r.POST("/stats", handler.GetDnaReportStats)
}

// IsMutant godoc
// @Summary Check if given dna is from a mutant
// @Description return true or false by dna input
// @Accept  json
// @Produce  json
// @Param body body models.IsMutantReqBody true "dna data"
// @Success 200
// @Failure 403 {object} apperrors.DeliveryError
// @Router /mutant [post]
// IsMutant handlr for http request to isMutant use case
func (h *DnaHandler) IsMutant(c *gin.Context) {
	var (
		req           models.IsMutantReqBody
		appError      *apperrors.AppError
		deliveryError *apperrors.DeliveryError
		isMutant      bool
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		deliveryError = &apperrors.DeliveryError{
			Err:         apperrors.ErrBadParamInput.Error(),
			Description: err.Error(),
		}
		c.JSON(http.StatusBadRequest, deliveryError)
		return
	}

	isMutant, appError = h.IsMutantUcase.Execute(req.Dna)

	if appError != nil {
		deliveryError = mappers.MapAppErrorToDeliveryError(appError)
		c.JSON(http.StatusBadRequest, deliveryError)
		return
	}

	if !isMutant {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	return
}

// GetDnaReportStats godoc
// @Summary returns db stats about dna
// @Description return db stats about dna
// @Accept  json
// @Produce  json
// @Success 200 {object} dna.DnaReportStats
// @Failure 400 {object} apperrors.DeliveryError
// @Router /stats [post]
// GetDnaReportStats handlr for http request to getDnaReportStats use case
func (h *DnaHandler) GetDnaReportStats(c *gin.Context) {
	var (
		appError      *apperrors.AppError
		deliveryError *apperrors.DeliveryError
	)

	reportResult, appError := h.GetDnaReportStatsUcase.Execute()

	if appError != nil {
		deliveryError = mappers.MapAppErrorToDeliveryError(appError)
		c.JSON(http.StatusBadRequest, deliveryError)
		return
	}

	c.JSON(200, reportResult)
	return
}
