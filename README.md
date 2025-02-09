# Tempo - A Simple Go Timer Package

`tempo` is a Go package for tracking and logging the execution time of processes. It supports optional configurations such as labels and thresholds to mark slow execution.

## Installation

You can install the package using `go get`:

```bash
go get github.com/barkhayot/tempo
```

## Usage

### Basic Timer

```go
package main

import (
	"time"
	"github.com/barkhayot/tempo"
)

func main() {
	timer := tempo.New()
	time.Sleep(2 * time.Second)
	timer.Stop()
}
```

### Timer with Label and Threshold

```go
package main

import (
	"time"
	"github.com/barkhayot/tempo"
)

func main() {
    timer := tempo.New(
        tempo.WithLabel("Process A"),
        tempo.WithThreshold(1*time.Second),
    )
    time.Sleep(3 * time.Second)
    timer.Stop()
}
```

### Run with Timeout

```go
package main

import (
	"time"
	"github.com/barkhayot/tempo"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    tempo.RunWithTimeout(ctx, func() {
        time.Sleep(3 * time.Second) // Simulate slow function
    }, tempo.WithLabel("Timeout Example"))
}
```