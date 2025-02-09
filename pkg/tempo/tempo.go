package tempo

import (
	"context"
	"log"
	"time"
)

type Tempo struct {
	startTime time.Time
	label     string
	threshold time.Duration
	logger    Logger
}

type Logger interface {
	Printf(format string, v ...interface{})
}

type Option func(*Tempo)

func WithLabel(label string) Option {
	return func(t *Tempo) {
		t.label = label
	}
}

func WithThreshold(threshold time.Duration) Option {
	return func(t *Tempo) {
		t.threshold = threshold
	}
}

func WithLogger(logger Logger) Option {
	return func(t *Tempo) {
		t.logger = logger
	}
}

func New(options ...Option) *Tempo {
	t := &Tempo{
		startTime: time.Now(),
		logger:    log.New(log.Writer(), "", log.LstdFlags), // default logger
	}

	for _, opt := range options {
		opt(t)
	}

	return t
}

func (t *Tempo) Stop() {
	elapsed := time.Since(t.startTime)
	if t.label == "" {
		t.logger.Printf("\u23f1\ufe0f  Process time: %s\n", elapsed)
	} else if t.threshold > 0 && elapsed > t.threshold {
		t.logger.Printf("[\u26a0\ufe0f  SLOW] %s took %s (threshold: %s)", t.label, elapsed, t.threshold)
	} else {
		t.logger.Printf("[\u2705 OK] %s  took %s", t.label, elapsed)
	}
}

func RunWithTimeout(ctx context.Context, fn func(), options ...Option) {
	t := New(options...)
	done := make(chan struct{})
	go func() {
		fn()
		close(done)
	}()

	select {
	case <-done:
		t.Stop()
	case <-ctx.Done():
		t.logger.Printf("[\u23f3 TIMEOUT] %s exceeded timeout", t.label)
	}
}
