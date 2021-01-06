package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/services"
)

// Resolver ...
type Resolver struct {
	AudiencesService     services.AudiencesService
	NotificationsService services.NotificationsService
	ExportsService       services.ExportsService
	SegmentsService      services.SegmentsService
}
