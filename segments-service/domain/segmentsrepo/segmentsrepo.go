package segmentsrepo

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
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
	segments, err := r.ListSegments(ctx)
	if err != nil {
		return nil, err
	}

	for _, segment := range segments {
		if segment.ID == id {
			return segment, nil
		}
	}

	return nil, domain.ErrEntityNotFound
}

// ListSegments ...
func (r Repository) ListSegments(ctx context.Context) ([]*domain.Segment, error) {
	file, err := os.Open("config/segments.json")
	if err != nil {
		return nil, err
	}

	segments := []*domain.Segment{}

	parser := json.NewDecoder(file)
	if err := parser.Decode(&segments); err != nil {
		return nil, err
	}

	return segments, nil
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
