Quick-start
-----------

```go
package main

import (
  "fmt"
  "time"

  "github.com/BradburyLab/go/metrics"
)

func main() {
  ms := metrics.New().Register(
    metrics.CPU(),
    metrics.GPU(),
    metrics.LoadAVG().Register(
      metrics.LoadAVG1(),
      metrics.LoadAVG5(),
      metrics.LoadAVG15(),
    ),
    metrics.MemVirt().Register(
      metrics.MemVirtTotal(),
      metrics.MemVirtAvailable(),
      metrics.MemVirtUsed(),
      metrics.MemVirtUsedPercent(),
      metrics.MemVirtFree(),
    ),
    metrics.MemSwap().Register(
      metrics.MemSwapTotal(),
      metrics.MemSwapFree(),
      metrics.MemSwapUsed(),
      metrics.MemSwapUsedPercent(),
    ),
    metrics.DisksIO().Register(
      metrics.DiskIO("sda").Register(
        metrics.DiskIOReadCount(),
        metrics.DiskIOWriteCount(),
        metrics.DiskIOReadBytes(),
        metrics.DiskIOWriteBytes(),
      ),
    ),
    metrics.DiskUsage("/").Register(
      metrics.DiskUsageTotal(),
      metrics.DiskUsageFree(),
      metrics.DiskUsageUsed(),
      metrics.DiskUsageUsedPercent(),
    ),
  )

  for {
    top := metrics.NewTop().Reduce(ms.Collect())
    fmt.Println(top)
    if !top.OK() {
      for _, err := range top.Errs() {
        fmt.Printf("%18s %s\n", err.Kind, err.Err.Error())
      }
    }

    time.Sleep(1 * time.Second)
  }
}
```