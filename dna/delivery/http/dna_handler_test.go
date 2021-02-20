package http_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	httpdelivery "mutant_detector/dna/delivery/http"
	"mutant_detector/dna/delivery/http/models"
	apperrors "mutant_detector/domain/app_errors"
	"mutant_detector/domain/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	dnamocks "mutant_detector/dna/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIsMutant(t *testing.T) {
	t.Run("Should return http 200 status code for mutant", func(t *testing.T) {
		router := gin.Default()
		mockIsmutantUCase := new(mocks.IsMutantUseCaseMock)
		rr := httptest.NewRecorder()
		mockedReqBody := models.IsMutantReqBody{
			Dna: dnamocks.DnaCorrectFormatInput,
		}

		mockIsmutantUCase.On("Execute", dnamocks.DnaCorrectFormatInput).Return(true, nil)

		httpdelivery.NewDnaHandler(router, mockIsmutantUCase)

		marshaledBody, err := json.Marshal(mockedReqBody)
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewBuffer(marshaledBody))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Should return 400 status code and bad-input-error if wrong body was sent", func(t *testing.T) {

		var errorResponse apperrors.DeliveryError
		router := gin.Default()
		mockIsmutantUCase := new(mocks.IsMutantUseCaseMock)
		rr := httptest.NewRecorder()
		mockedReqBody := map[string]interface{}{
			"test": dnamocks.DnaCorrectFormatInput,
		}

		mockIsmutantUCase.On("Execute", dnamocks.DnaCorrectFormatInput).Return(true, nil)

		httpdelivery.NewDnaHandler(router, mockIsmutantUCase)

		marshaledBody, err := json.Marshal(mockedReqBody)
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewBuffer(marshaledBody))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		mockIsmutantUCase.AssertNotCalled(t, "Execute", mock.Anything)

		err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)
		fmt.Println(errorResponse)
		responseError := errors.New(errorResponse.Err)

		assert.Equal(t, responseError, apperrors.ErrBadParamInput)
		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("Should return 400 status code and bad-input-error if isMutant ucase returns an error", func(t *testing.T) {

		var errorResponse apperrors.DeliveryError
		router := gin.Default()
		mockIsmutantUCase := new(mocks.IsMutantUseCaseMock)
		rr := httptest.NewRecorder()
		mockedReqBody := models.IsMutantReqBody{
			Dna: dnamocks.DnaCorrectFormatInput,
		}

		expectedAppError := &apperrors.AppError{
			Err:         apperrors.ErrBadParamInput,
			Description: "some",
		}

		mockIsmutantUCase.On("Execute", dnamocks.DnaCorrectFormatInput).Return(false, expectedAppError)

		httpdelivery.NewDnaHandler(router, mockIsmutantUCase)

		marshaledBody, err := json.Marshal(mockedReqBody)
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewBuffer(marshaledBody))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)
		fmt.Println(errorResponse)
		responseError := errors.New(errorResponse.Err)

		assert.Equal(t, responseError, expectedAppError.Err)
		assert.Equal(t, errorResponse.Description, expectedAppError.Description)
		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

	t.Run("Should return 403 status code for human", func(t *testing.T) {
		router := gin.Default()
		mockIsmutantUCase := new(mocks.IsMutantUseCaseMock)
		rr := httptest.NewRecorder()
		mockedReqBody := models.IsMutantReqBody{
			Dna: dnamocks.DnaCorrectFormatInput,
		}

		mockIsmutantUCase.On("Execute", dnamocks.DnaCorrectFormatInput).Return(false, nil)

		httpdelivery.NewDnaHandler(router, mockIsmutantUCase)

		marshaledBody, err := json.Marshal(mockedReqBody)
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewBuffer(marshaledBody))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusForbidden, rr.Code)
	})
}
