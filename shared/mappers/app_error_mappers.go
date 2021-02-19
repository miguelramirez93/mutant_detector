package mappers

import "mutant_detector/domain"

// MapAppErrorToDeliveryError map given appError to deliveryError
func MapAppErrorToDeliveryError(appError *domain.AppError) *domain.DeliveryError {
	return &domain.DeliveryError{
		Err:         appError.Err.Error(),
		Description: appError.Description,
	}
}
