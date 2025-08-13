package apiKey

import (
	"context"
	"errors"

	configs "rate-limiter/configs"
	"rate-limiter/internal/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

//go:generate mockgen -source=api_key.service.go -destination=api_key.service_mock.go -package=api_key
type Repository interface {
	CreateAPIKey(ctx context.Context, params db.CreateAPIKeyParams) (db.ApiKey, error)
	GetAPIKey(ctx context.Context, key uuid.UUID) (db.ApiKey, error)
	UpdateAPIKey(ctx context.Context, params db.UpdateAPIKeyParams) (db.ApiKey, error)
	DeleteAPIKey(ctx context.Context, id int32) error
}

type ApiKeyService struct {
	config     *configs.Config
	repository Repository
}

func NewApiKeyService(config *configs.Config, repository Repository) *ApiKeyService {
	return &ApiKeyService{config: config, repository: repository}
}

type CreateAPIKeyParams struct {
	AccountID   uuid.UUID
	Name        string
	Description *string
	Key         uuid.UUID
	Status      db.Status
	CreatedBy   *string
	Product     db.Product
	PlanID      int32
}

type UpdateAPIKeyParams struct {
	Name        string
	Description *string
	Status      db.Status
}

// CreateAPIKey creates a new API key
func (s *ApiKeyService) CreateAPIKey(ctx context.Context, params *CreateAPIKeyParams) (*db.ApiKey, error) {
	// Check if API key already exists
	_, err := s.repository.GetAPIKey(ctx, params.Key)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			s.config.Logger.Error("Failed to check if API key exists", "error", err)
			return nil, err
		}
	} else {
		s.config.Logger.Error("API key already exists")
		return nil, errors.New("API key already exists")
	}

	createParams := MapCreateAPIKeyParamsToDbCreateAPIKeyParams(params)

	createdAPIKey, err := s.repository.CreateAPIKey(ctx, createParams)
	if err != nil {
		s.config.Logger.Error("Failed to create API key", "error", err)
		return nil, err
	}

	return &createdAPIKey, nil
}

// GetAPIKey gets an API key by key
func (s *ApiKeyService) GetAPIKey(ctx context.Context, key uuid.UUID) (*db.ApiKey, error) {
	apiKey, err := s.repository.GetAPIKey(ctx, key)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrAPIKeyNotFound{}
		}
		s.config.Logger.Error("Failed to get API key", "error", err)
		return nil, err
	}
	return &apiKey, nil
}

// UpdateAPIKey updates an API key
func (s *ApiKeyService) UpdateAPIKey(ctx context.Context, params *UpdateAPIKeyParams) (*db.ApiKey, error) {
	updateParams := MapUpdateAPIKeyParamsToDbUpdateAPIKeyParams(params)

	updatedAPIKey, err := s.repository.UpdateAPIKey(ctx, updateParams)
	if err != nil {
		s.config.Logger.Error("Failed to update API key", "error", err)
		return nil, err
	}

	return &updatedAPIKey, nil
}

// DeleteAPIKey deletes an API key by ID
func (s *ApiKeyService) DeleteAPIKey(ctx context.Context, id int32) error {
	err := s.repository.DeleteAPIKey(ctx, id)
	if err != nil {
		s.config.Logger.Error("Failed to delete API key", "error", err)
		return err
	}

	return nil
}
