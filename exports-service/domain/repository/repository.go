package repository

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/lcnascimento/istio-knative-poc/exports-service/domain"
)

// ServiceInput ...
type ServiceInput struct {
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	return &Service{in: in}, nil
}

// GetExport ...
func (s Service) GetExport(ctx context.Context, id string) (*domain.Export, error) {
	log.Printf("Fetch export %s from database", id)

	exports, err := s.ListExports(ctx)
	if err != nil {
		return nil, err
	}

	for _, export := range exports {
		if export.ID == id {
			return export, nil
		}
	}

	return nil, domain.ErrEntityNotFound
}

// ListExports ...
func (s Service) ListExports(ctx context.Context) ([]*domain.Export, error) {
	log.Printf("Fetch exports from database")

	file, err := os.Open("config/exports.json")
	if err != nil {
		return nil, err
	}

	exports := []*domain.Export{}

	parser := json.NewDecoder(file)
	if err := parser.Decode(&exports); err != nil {
		return nil, err
	}

	return exports, nil
}
