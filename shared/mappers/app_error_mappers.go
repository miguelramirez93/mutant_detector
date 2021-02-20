package mappers

import (
	apperrors "mutant_detector/domain/app_errors"
)

// MapAppErrorToDeliveryError map given appError to deliveryError
func MapAppErrorToDeliveryError(appError *apperrors.AppError) *apperrors.DeliveryError {
	return &apperrors.DeliveryError{
		Err:         appError.Err.Error(),
		Description: appError.Description,
	}
}
