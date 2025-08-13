package plan

import (
	"context"
	"errors"

	configs "rate-limiter/configs"
	"rate-limiter/internal/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

//go:generate mockgen -source=plan.service.go -destination=plan.service_mock.go -package=plan
type Repository interface {
	CreatePlan(ctx context.Context, params db.CreatePlanParams) (db.Plan, error)
	GetPlanByName(ctx context.Context, name string) (db.Plan, error)
	UpdatePlan(ctx context.Context, params db.UpdatePlanParams) (db.Plan, error)
	DeletePlan(ctx context.Context, name string) error
}

type PlanService struct {
	config     *configs.Config
	repository Repository
}

func NewPlanService(config *configs.Config, repository Repository) *PlanService {
	return &PlanService{config: config, repository: repository}
}

type CreatePlanParams struct {
	AccountID uuid.UUID
	Name      string
	RateLimit int32
	BurstSize int32
	Algorithm db.Algorithm
}

type UpdatePlanParams struct {
	Name      string
	RateLimit int32
	BurstSize int32
	Algorithm db.Algorithm
}

// CreatePlan creates a new plan
func (s *PlanService) CreatePlan(ctx context.Context, params *CreatePlanParams) (*db.Plan, error) {
	// Check if plan already exists
	_, err := s.repository.GetPlanByName(ctx, params.Name)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			s.config.Logger.Error("Failed to check if plan exists", "error", err)
			return nil, err
		}
	} else {
		s.config.Logger.Error("Plan already exists")
		return nil, errors.New("plan already exists")
	}

	createParams := MapCreatePlanParamsToDbCreatePlanParams(params)

	createdPlan, err := s.repository.CreatePlan(ctx, createParams)
	if err != nil {
		s.config.Logger.Error("Failed to create plan", "error", err)
		return nil, err
	}

	return &createdPlan, nil
}

// GetPlanByName gets a plan by name
func (s *PlanService) GetPlanByName(ctx context.Context, name string) (*db.Plan, error) {
	plan, err := s.repository.GetPlanByName(ctx, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrPlanNotFound{}
		}
		s.config.Logger.Error("Failed to get plan by name", "error", err)
		return nil, err
	}
	return &plan, nil
}

// UpdatePlan updates a plan
func (s *PlanService) UpdatePlan(ctx context.Context, params *UpdatePlanParams) (*db.Plan, error) {
	updateParams := MapUpdatePlanParamsToDbUpdatePlanParams(params)

	updatedPlan, err := s.repository.UpdatePlan(ctx, updateParams)
	if err != nil {
		s.config.Logger.Error("Failed to update plan", "error", err)
		return nil, err
	}

	return &updatedPlan, nil
}

// DeletePlan deletes a plan by name
func (s *PlanService) DeletePlan(ctx context.Context, name string) error {
	err := s.repository.DeletePlan(ctx, name)
	if err != nil {
		s.config.Logger.Error("Failed to delete plan", "error", err)
		return err
	}

	return nil
}
