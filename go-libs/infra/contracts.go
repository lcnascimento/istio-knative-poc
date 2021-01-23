package golibs

import (
	"context"
)

// LogProvider provides functions to log information with
// different visibility levels in the following order
// from least important to more important:
//
// - Debug
// - Info
// - Warning
// - Error
// - Critical
type LogProvider interface {
	Debug(ctx context.Context, msg string, args ...interface{})
	Info(ctx context.Context, msg string, args ...interface{})
	Warning(ctx context.Context, msg string, args ...interface{})
	Error(ctx context.Context, err error)
	Critical(ctx context.Context, err error)
}
