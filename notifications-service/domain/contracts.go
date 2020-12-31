package domain

import (
	"context"
	"fmt"
)

// ErrEntityNotFound ...
var ErrEntityNotFound = fmt.Errorf("entity not found")

// ErrNotImplemented ...
var ErrNotImplemented = fmt.Errorf("method not implemented yet")

// Channel ...
type Channel string

const (
	// EmailChannel ...
	EmailChannel Channel = "email"
	// SMSChannel ...
	SMSChannel Channel = "sms"
	// BrowserPushChannel ...
	BrowserPushChannel Channel = "browser"
)

// NotificationsRepository ...
type NotificationsRepository interface {
	GetNotification(ctx context.Context, id string) (*Notification, error)
	ListNotifications(ctx context.Context) ([]*Notification, error)
}

// NotificationsEnqueuer ...
type NotificationsEnqueuer interface {
	EnqueueNotification(ctx context.Context, id string) error
}

// NotificationsSender ...
type NotificationsSender interface {
	SendNotification(ctx context.Context, id string) error
}

// NotificationProvider ...
type NotificationProvider interface {
	SendNotification(ctx context.Context, notif Notification, ch chan []*User) (chan bool, chan error)
}

// SegmentsService ...
type SegmentsService interface {
	GetSegmentUsers(ctx context.Context, id string) (chan []*User, chan error)
}
