package sender

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"

	infra "github.com/lcnascimento/istio-knative-poc/go-libs/infra"

	"github.com/lcnascimento/istio-knative-poc/audiences-service/domain"
)

// ServiceInput ...
type ServiceInput struct {
	Tracer  trace.Tracer
	Logger  infra.LogProvider
	Repo    domain.AudiencesRepository
	Exports domain.ExportsService
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	if in.Logger == nil {
		return nil, fmt.Errorf("Missing required dependency: Logger")
	}

	if in.Tracer == nil {
		return nil, fmt.Errorf("Missing required dependency: Tracer")
	}

	return &Service{in: in}, nil
}

// SendAudience ...
func (s Service) SendAudience(ctx context.Context, id string, exportID string) error {
	ctx, span := s.in.Tracer.Start(ctx, "domain.sender.SendAudience")
	defer span.End()

	s.in.Logger.Info(ctx, "Sending audience %s", id)

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

	s.in.Logger.Info(ctx, "Audience %s sent successfuly", id)
	return nil
}

func (s Service) downloadExport(ctx context.Context, url string) (string, error) {
	ctx, span := s.in.Tracer.Start(ctx, "domain.sender.downloadExport")
	defer span.End()

	s.in.Logger.Info(ctx, "Downloading export in %s", url)

	return "fake/path/to/export/file", nil
}

func (s Service) doDiff(ctx context.Context, oldPath, newPath string) error {
	ctx, span := s.in.Tracer.Start(ctx, "domain.sender.doDiff")
	defer span.End()

	s.in.Logger.Info(ctx, "Applying diff between old and new exports")

	return nil
}
