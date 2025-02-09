package main

import (
	"context"
	"fmt"
	"time"

	"github.com/barkhayot/tempo/pkg/tempo"
)

func UsageWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tempo.RunWithTimeout(ctx, func() {
		time.Sleep(3 * time.Second)
	}, tempo.WithLabel("Timeout Example"))

	fmt.Println("Timeout example completed")
}
