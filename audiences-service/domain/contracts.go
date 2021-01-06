package domain

import (
	"context"
	"fmt"
)

// ErrEntityNotFound ...
var ErrEntityNotFound = fmt.Errorf("entity not found")

// ErrNotImplemented ...
var ErrNotImplemented = fmt.Errorf("method not implemented yet")

// AudiencesRepository ...
type AudiencesRepository interface {
	GetAudience(ctx context.Context, id string) (*Audience, error)
	ListAudiences(ctx context.Context) ([]*Audience, error)
}

// AudienceSendingEnqueuer ...
type AudienceSendingEnqueuer interface {
	EnqueueAudienceSending(ctx context.Context, id string) error
}

// AudiencesSender ...
type AudiencesSender interface {
	SendAudience(ctx context.Context, id string) error
}

// ExportsService ...
type ExportsService interface {
	CreateExport(ctx context.Context, expo Export) error
	GetExportDownloadURL(ctx context.Context, exportID string) (string, error)
}
