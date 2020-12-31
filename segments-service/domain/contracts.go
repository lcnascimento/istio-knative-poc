package domain

import (
	"context"
	"fmt"
)

// ErrEntityNotFound ...
var ErrEntityNotFound = fmt.Errorf("entity not found")

// ErrNotImplemented ...
var ErrNotImplemented = fmt.Errorf("method not implemented yet")

// SegmentsRepository ...
type SegmentsRepository interface {
	FindSegment(ctx context.Context, id string) (*Segment, error)
	ListSegments(ctx context.Context) ([]*Segment, error)
	GetSegmentUsers(ctx context.Context, id string, options GetSegmentUsersOptions) (chan []*User, error)
}

// GetSegmentUsersOptions ...
type GetSegmentUsersOptions struct {
	BulkSize int
}
