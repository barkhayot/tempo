package main

import (
	"time"

	"github.com/barkhayot/tempo/pkg/tempo"
)

func BasicUsage() {
	timer := tempo.New()
	time.Sleep(10 * time.Second)
	timer.Stop()
}
