package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/generated"
	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/model"
)

func (r *exportResolver) Segment(ctx context.Context, obj *model.Export) (*model.Segment, error) {
	return r.SegmentsService.GetSegment(ctx, obj.Segment.ID)
}

func (r *mutationResolver) SendNotification(ctx context.Context, notificationID string) (*string, error) {
	return nil, r.NotificationsService.SendNotification(ctx, notificationID)
}

func (r *mutationResolver) CreateExport(ctx context.Context, input model.NewExport) (*model.Export, error) {
	return r.ExportsService.CreateExport(ctx, input)
}

func (r *notificationResolver) Segment(ctx context.Context, obj *model.Notification) (*model.Segment, error) {
	return r.SegmentsService.GetSegment(ctx, obj.Segment.ID)
}

func (r *queryResolver) Exports(ctx context.Context) ([]*model.Export, error) {
	return r.ExportsService.ListExports(ctx)
}

func (r *queryResolver) Export(ctx context.Context, id string) (*model.Export, error) {
	return r.ExportsService.GetExport(ctx, id)
}

func (r *queryResolver) Notifications(ctx context.Context) ([]*model.Notification, error) {
	return r.NotificationsService.ListNotifications(ctx)
}

func (r *queryResolver) Notification(ctx context.Context, id string) (*model.Notification, error) {
	return r.NotificationsService.GetNotification(ctx, id)
}

func (r *queryResolver) Segments(ctx context.Context) ([]*model.Segment, error) {
	return r.SegmentsService.ListSegments(ctx)
}

func (r *queryResolver) Segment(ctx context.Context, id string) (*model.Segment, error) {
	return r.SegmentsService.GetSegment(ctx, id)
}

// Export returns generated.ExportResolver implementation.
func (r *Resolver) Export() generated.ExportResolver { return &exportResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Notification returns generated.NotificationResolver implementation.
func (r *Resolver) Notification() generated.NotificationResolver { return &notificationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type exportResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type notificationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
