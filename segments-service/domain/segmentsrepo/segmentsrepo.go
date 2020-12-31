package segmentsrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/lcnascimento/istio-knative-poc/segments-service/domain"
)

// RepositoryInput ...
type RepositoryInput struct {
	NumberOfUsersInSegments int
	NetworkDelay            time.Duration
}

// Repository ...
type Repository struct {
	in RepositoryInput
}

// NewRepository ...
func NewRepository(in RepositoryInput) (*Repository, error) {
	if in.NumberOfUsersInSegments == 0 {
		in.NumberOfUsersInSegments = 100
	}

	return &Repository{in: in}, nil
}

// FindSegment ...
func (r Repository) FindSegment(ctx context.Context, id string) (*domain.Segment, error) {
	return nil, domain.ErrNotImplemented
}

// ListSegments ...
func (r Repository) ListSegments(ctx context.Context) ([]*domain.Segment, error) {
	return nil, domain.ErrNotImplemented
}

// GetSegmentUsers ...
func (r Repository) GetSegmentUsers(ctx context.Context, id string, options domain.GetSegmentUsersOptions) (chan []*domain.User, error) {
	ch := make(chan []*domain.User)

	go func() {
		numBulks := r.in.NumberOfUsersInSegments / options.BulkSize

		for i := 0; i < numBulks; i++ {
			bulk := []*domain.User{}

			time.Sleep(r.in.NetworkDelay)
			for j := 0; j < options.BulkSize; j++ {
				bulk = append(bulk, &domain.User{
					Reference: fmt.Sprintf("%d_%d", i, j),
					AppKey:    "Random AppKey",
					Name:      fmt.Sprintf("Anonymous %d_%d", i, j),
					Email:     fmt.Sprintf("random+%d_%d@email.com", i, j),
				})
			}

			ch <- bulk
		}

		close(ch)
	}()

	return ch, nil
}
