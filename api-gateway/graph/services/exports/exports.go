package exports

import (
	"context"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/model"
	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/services"
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

// ListExports ...
func (s Service) ListExports(ctx context.Context) ([]*model.Export, error) {
	return nil, services.ErrNotImplemented
}

// GetExport ...
func (s Service) GetExport(ctx context.Context, id string) (*model.Export, error) {
	return nil, services.ErrNotImplemented
}

// CreateExport ...
func (s Service) CreateExport(ctx context.Context, in model.NewExport) (*model.Export, error) {
	return nil, services.ErrNotImplemented
}
