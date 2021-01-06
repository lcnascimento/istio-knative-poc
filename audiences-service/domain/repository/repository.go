package repository

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/lcnascimento/istio-knative-poc/audiences-service/domain"
)

// ServiceInput ...
type ServiceInput struct{}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	return &Service{in: in}, nil
}

// GetAudience ...
func (s Service) GetAudience(ctx context.Context, id string) (*domain.Audience, error) {
	log.Printf("Fetch audience %s from database", id)

	audiences, err := s.ListAudiences(ctx)
	if err != nil {
		return nil, err
	}

	for _, audience := range audiences {
		if audience.ID == id {
			return audience, nil
		}
	}

	return nil, domain.ErrEntityNotFound
}

// ListAudiences ...
func (s Service) ListAudiences(ctx context.Context) ([]*domain.Audience, error) {
	log.Printf("Fetch audiences from database")

	file, err := os.Open("config/audiences.json")
	if err != nil {
		return nil, err
	}

	audiences := []*domain.Audience{}

	parser := json.NewDecoder(file)
	if err := parser.Decode(&audiences); err != nil {
		return nil, err
	}

	return audiences, nil
}
