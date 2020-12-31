package exportator

import (
	"context"
	"fmt"
	"log"

	"github.com/lcnascimento/istio-knative-poc/exports-service/domain"
	"golang.org/x/sync/errgroup"
)

// ServiceInput ...
type ServiceInput struct {
	Repo     domain.ExportsRepository
	Segments domain.SegmentsService
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	if in.Repo == nil {
		return nil, fmt.Errorf("Missing ExportsRepository dependency")
	}

	if in.Segments == nil {
		return nil, fmt.Errorf("Missing SegmentsService dependency")
	}

	return &Service{in: in}, nil
}

// Export ...
func (s Service) Export(ctx context.Context, id string) error {
	log.Printf("Starting export %s", id)

	expo, err := s.in.Repo.GetExport(ctx, id)
	if err != nil {
		return err
	}

	g, gctx := errgroup.WithContext(ctx)

	ch, errCh := s.in.Segments.GetSegmentUsers(gctx, expo.SegmentID)

	g.Go(func() error {
		for bulk := range ch {
			log.Printf("Exporting %d users to CSV file", len(bulk))
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

	log.Printf("Export %s finished successfuly", id)
	return nil
}
