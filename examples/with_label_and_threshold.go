package main

import (
	"fmt"
	"time"

	"github.com/barkhayot/tempo/pkg/tempo"
)

func UsageWithOptions() {
	timer := tempo.New(
		tempo.WithLabel("Process A"),
		tempo.WithThreshold(5*time.Second),
	)
	time.Sleep(3 * time.Second)
	timer.Stop()

	fmt.Println("Label and Threshold example completed")
}
