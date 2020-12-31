package domain

import (
	"context"
	"fmt"
)

// ErrEntityNotFound ...
var ErrEntityNotFound = fmt.Errorf("entity not found")

// ErrNotImplemented ...
var ErrNotImplemented = fmt.Errorf("method not implemented yet")

// ExportsRepository ...
type ExportsRepository interface {
	GetExport(ctx context.Context, id string) (*Export, error)
	ListExports(ctx context.Context) ([]*Export, error)
}

// ExportEnqueuer ...
type ExportEnqueuer interface {
	EnqueueExport(ctx context.Context, id string) error
}

// Exportator ...
type Exportator interface {
	Export(ctx context.Context, id string) error
}

// SegmentsService ...
type SegmentsService interface {
	GetSegmentUsers(ctx context.Context, id string) (chan []*User, chan error)
}
