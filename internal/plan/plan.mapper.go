package plan

import (
	"rate-limiter/internal/db"
)

// MapCreatePlanParamsToDbCreatePlanParams converts service params to database params
func MapCreatePlanParamsToDbCreatePlanParams(params *CreatePlanParams) db.CreatePlanParams {
	return db.CreatePlanParams{
		AccountID: params.AccountID,
		Name:      params.Name,
		RateLimit: params.RateLimit,
		BurstSize: params.BurstSize,
		Algorithm: params.Algorithm,
	}
}

// MapUpdatePlanParamsToDbUpdatePlanParams converts service update params to database params
func MapUpdatePlanParamsToDbUpdatePlanParams(params *UpdatePlanParams) db.UpdatePlanParams {
	return db.UpdatePlanParams{
		Name:      params.Name,
		RateLimit: params.RateLimit,
		BurstSize: params.BurstSize,
		Algorithm: params.Algorithm,
	}
}
