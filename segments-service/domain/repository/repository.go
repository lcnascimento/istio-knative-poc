package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"go.opentelemetry.io/otel/trace"

	infra "github.com/lcnascimento/istio-knative-poc/go-libs/infra"

	"github.com/lcnascimento/istio-knative-poc/segments-service/domain"
)

// ServiceInput ...
type ServiceInput struct {
	NumberOfUsersInSegments int
	NetworkDelay            time.Duration

	Tracer trace.Tracer
	Logger infra.LogProvider
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	if in.Tracer == nil {
		return nil, fmt.Errorf("Missing required dependency: Tracer")
	}

	if in.Logger == nil {
		return nil, fmt.Errorf("Missing required dependency: Logger")
	}

	if in.NumberOfUsersInSegments == 0 {
		in.NumberOfUsersInSegments = 100
	}

	return &Service{in: in}, nil
}

// FindSegment ...
func (r Service) FindSegment(ctx context.Context, id string) (*domain.Segment, error) {
	ctx, span := r.in.Tracer.Start(ctx, "domain.repository.FindSegment")
	defer span.End()

	r.in.Logger.Info(ctx, "Fetching segment %s from database", id)

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
func (r Service) ListSegments(ctx context.Context) ([]*domain.Segment, error) {
	ctx, span := r.in.Tracer.Start(ctx, "domain.repository.ListSegments")
	defer span.End()

	r.in.Logger.Info(ctx, "Fetching segments from database")
	time.Sleep(r.in.NetworkDelay)

	file, err := os.Open("config/segments.json")
	if err != nil {
		r.in.Logger.Error(ctx, fmt.Errorf("could not open config/segments.json file: %s", err.Error()))
		return nil, err
	}

	segments := []*domain.Segment{}

	parser := json.NewDecoder(file)
	if err := parser.Decode(&segments); err != nil {
		r.in.Logger.Error(ctx, fmt.Errorf("could not decode segments json file: %s", err.Error()))
		return nil, err
	}

	return segments, nil
}

// GetSegmentUsers ...
func (r Service) GetSegmentUsers(ctx context.Context, id string, options domain.GetSegmentUsersOptions) (chan []*domain.User, chan error) {
	ctx, span := r.in.Tracer.Start(ctx, "domain.repository.GetSegmentUsers")

	ch := make(chan []*domain.User)
	errCh := make(chan error)

	go func() {
		defer span.End()
		defer close(ch)
		defer close(errCh)

		_, err := r.FindSegment(ctx, id)
		if err != nil {
			errCh <- err
			return
		}

		numBulks := r.in.NumberOfUsersInSegments / options.BulkSize

		for i := 0; i < numBulks; i++ {
			bulk := []*domain.User{}

			r.in.Logger.Info(ctx, "Fetching users from database")
			time.Sleep(r.in.NetworkDelay)
			for j := 0; j < options.BulkSize; j++ {
				bulk = append(bulk, &domain.User{
					Reference: fmt.Sprintf("%d_%d", i, j),
					AppKey:    "Random AppKey",
					Name:      fmt.Sprintf("Anonymous %d_%d", i, j),
					Email:     fmt.Sprintf("random+%d_%d@email.com", i, j),
				})
			}

			select {
			case <-ctx.Done():
				return
			case ch <- bulk:
			}
		}

		r.in.Logger.Info(ctx, "All users extracted successfuly")
	}()

	return ch, errCh
}
