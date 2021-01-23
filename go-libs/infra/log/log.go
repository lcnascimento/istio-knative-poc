package log

import (
	"context"
	"fmt"
	"log"

	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/errors"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
)

// Level ...
type Level string

var (
	// DebugLevel ...
	DebugLevel Level = "DEBUG"
	// InfoLevel ...
	InfoLevel Level = "INFO"
	// WarningLevel ...
	WarningLevel Level = "WARNING"
	// ErrorLevel ...
	ErrorLevel Level = "ERROR"
	// CriticalLevel ...
	CriticalLevel Level = "CRITICAL"
)

var levelPriority = map[Level]int{
	CriticalLevel: 1,
	ErrorLevel:    2,
	WarningLevel:  3,
	InfoLevel:     4,
	DebugLevel:    5,
}

// ClientInput ...
type ClientInput struct {
	Level Level
}

// Client ...
type Client struct {
	in ClientInput
}

// NewClient ...
func NewClient(in ClientInput) (*Client, error) {
	if _, ok := levelPriority[in.Level]; !ok {
		in.Level = InfoLevel
	}

	return &Client{in: in}, nil
}

// Debug ...
func (c Client) Debug(ctx context.Context, msg string, args ...interface{}) {
	if levelPriority[c.in.Level] >= 5 {
		c.print(ctx, fmt.Sprintf(msg, args...), "debug")
	}
}

// Info ...
func (c Client) Info(ctx context.Context, msg string, args ...interface{}) {
	if levelPriority[c.in.Level] >= 4 {
		c.print(ctx, fmt.Sprintf(msg, args...), "info")
	}
}

// Warning ...
func (c Client) Warning(ctx context.Context, msg string, args ...interface{}) {
	if levelPriority[c.in.Level] >= 3 {
		c.print(ctx, fmt.Sprintf(msg, args...), "warning")
	}
}

// Error ...
func (c Client) Error(ctx context.Context, err error) {
	if levelPriority[c.in.Level] >= 2 {
		c.printError(ctx, err, "error")
	}
}

// Critical ...
func (c Client) Critical(ctx context.Context, err error) {
	if levelPriority[c.in.Level] >= 1 {
		c.printError(ctx, err, "critical")
	}
}

func (c Client) print(ctx context.Context, msg string, level string) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(
		"log",
		trace.WithAttributes(label.String("log.level", level)),
		trace.WithAttributes(label.String("log.message", msg)),
	)
	log.Println(msg)
}

func (c Client) printError(ctx context.Context, err error, level string) {
	span := trace.SpanFromContext(ctx)
	span.RecordError(
		err,
		trace.WithAttributes(label.String("log.level", level)),
		trace.WithAttributes(label.String("error.kind", errors.Kind(err).String())),
		trace.WithAttributes(label.String("error.severity", errors.Severity(err).String())),
	)
	log.Println(err.Error())
}
