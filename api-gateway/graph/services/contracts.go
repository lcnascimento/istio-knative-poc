package services

import (
	"context"
	"fmt"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/model"
)

// ErrNotImplemented ...
var ErrNotImplemented = fmt.Errorf("method not implemented yet")

// ExportsService ...
type ExportsService interface {
	ListExports(ctx context.Context) ([]*model.Export, error)
	GetExport(ctx context.Context, id string) (*model.Export, error)
	CreateExport(ctx context.Context, in model.NewExport) (*model.Export, error)
}

// NotificationsService ...
type NotificationsService interface {
	ListNotifications(ctx context.Context) ([]*model.Notification, error)
	GetNotification(ctx context.Context, id string) (*model.Notification, error)
	SendNotification(ctx context.Context, id string) error
}

// SegmentsService ...
type SegmentsService interface {
	ListSegments(ctx context.Context) ([]*model.Segment, error)
	GetSegment(ctx context.Context, id string) (*model.Segment, error)
}
