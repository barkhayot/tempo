package test

import (
	"context"
	"testing"
	"time"

	"github.com/barkhayot/tempo/pkg/tempo"
)

func TestTempo(t *testing.T) {
	t.Run("Test Basic Tempo", func(t *testing.T) {
		timer := tempo.New(
			tempo.WithLabel("Basic Test"),
		)
		time.Sleep(1 * time.Second)
		timer.Stop()
	})

	t.Run("Test Tempo with Threshold", func(t *testing.T) {
		timer := tempo.New(
			tempo.WithLabel("Slow Test"),
			tempo.WithThreshold(500*time.Millisecond),
		)
		time.Sleep(1 * time.Second)
		timer.Stop()
	})

	t.Run("Test Timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		tempo.RunWithTimeout(ctx, func() {
			time.Sleep(1 * time.Second)
		}, tempo.WithLabel("Timeout Test"))
	})
}
