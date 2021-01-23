package exportator

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"

	infra "github.com/lcnascimento/istio-knative-poc/go-libs/infra"

	"github.com/lcnascimento/istio-knative-poc/exports-service/domain"
)

// ServiceInput ...
type ServiceInput struct {
	Tracer   trace.Tracer
	Logger   infra.LogProvider
	Repo     domain.ExportsRepository
	Segments domain.SegmentsService
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

	if in.Repo == nil {
		return nil, fmt.Errorf("Missing required dependency: Repo")
	}

	if in.Segments == nil {
		return nil, fmt.Errorf("Missing required dependency: Segments")
	}

	return &Service{in: in}, nil
}

// Export ...
func (s Service) Export(ctx context.Context, id string) error {
	ctx, span := s.in.Tracer.Start(ctx, "domain.exportator.Export")
	defer span.End()

	s.in.Logger.Info(ctx, "Starting export %s", id)

	expo, err := s.in.Repo.GetExport(ctx, id)
	if err != nil {
		return err
	}

	g, gctx := errgroup.WithContext(ctx)

	ch, errCh := s.in.Segments.GetSegmentUsers(gctx, expo.SegmentID)

	g.Go(func() error {
		for bulk := range ch {
			s.in.Logger.Info(ctx, "Exporting %d users to CSV file", len(bulk))
		}

		return nil
	})

	g.Go(func() error {
		for err := range errCh {
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	s.in.Logger.Info(ctx, "Export %s finished successfuly", id)
	return nil
}
