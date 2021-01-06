package sender

import (
	"context"
	"log"

	"github.com/lcnascimento/istio-knative-poc/audiences-service/domain"
	"golang.org/x/sync/errgroup"
)

// ServiceInput ...
type ServiceInput struct {
	Repo    domain.AudiencesRepository
	Exports domain.ExportsService
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	return &Service{in: in}, nil
}

// SendAudience ...
func (s Service) SendAudience(ctx context.Context, id string, exportID string) error {
	log.Printf("Sending audience %s", id)

	aud, err := s.in.Repo.GetAudience(ctx, id)
	if err != nil {
		return err
	}

	var oldExportPath, newExportPath string

	g, gctx := errgroup.WithContext(ctx)

	if aud.LastExportID != "" {
		url, innerErr := s.in.Exports.GetExportDownloadURL(gctx, exportID)
		if innerErr != nil {
			return innerErr
		}

		oldExportPath, innerErr = s.downloadExport(gctx, url)
		return innerErr
	}

	g.Go(func() error {
		url, innerErr := s.in.Exports.GetExportDownloadURL(gctx, exportID)
		if innerErr != nil {
			return innerErr
		}

		newExportPath, innerErr = s.downloadExport(gctx, url)
		return innerErr
	})

	if err := g.Wait(); err != nil {
		return err
	}

	if err := s.doDiff(ctx, oldExportPath, newExportPath); err != nil {
		return err
	}

	log.Printf("Audience %s sent successfuly", id)
	return nil
}

func (s Service) downloadExport(ctx context.Context, url string) (string, error) {
	log.Printf("Downloading export in %s", url)

	return "fake/path/to/export/file", nil
}

func (s Service) doDiff(ctx context.Context, oldPath, newPath string) error {
	log.Println("Applying diff between old and new exports")

	return nil
}
