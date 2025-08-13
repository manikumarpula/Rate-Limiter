package apiKey

import (
	"rate-limiter/internal/db"
)

// MapCreateAPIKeyParamsToDbCreateAPIKeyParams converts service params to database params
func MapCreateAPIKeyParamsToDbCreateAPIKeyParams(params *CreateAPIKeyParams) db.CreateAPIKeyParams {
	return db.CreateAPIKeyParams{
		AccountID:   params.AccountID,
		Name:        params.Name,
		Description: params.Description,
		Key:         params.Key,
		Status:      params.Status,
		CreatedBy:   params.CreatedBy,
		Product:     params.Product,
		PlanID:      params.PlanID,
	}
}

// MapUpdateAPIKeyParamsToDbUpdateAPIKeyParams converts service update params to database params
func MapUpdateAPIKeyParamsToDbUpdateAPIKeyParams(params *UpdateAPIKeyParams) db.UpdateAPIKeyParams {
	return db.UpdateAPIKeyParams{
		Name:        params.Name,
		Description: params.Description,
		Status:      params.Status,
	}
}
