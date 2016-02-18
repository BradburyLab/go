package metrics

import (
	psutil "github.com/shirou/gopsutil/mem"
)

type MetricMemSwap interface {
	Collect(*psutil.SwapMemoryStat, error) Result
}

type memSwap struct {
	metrics []MetricMemSwap
}

func (it *memSwap) Register(metrics ...MetricMemSwap) *memSwap {
	for _, metric := range metrics {
		it.Append(metric)
	}
	return it
}

func (it *memSwap) Append(metric MetricMemSwap) *memSwap {
	it.metrics = append(it.metrics, metric)
	return it
}

func (it *memSwap) Len() int { return len(it.metrics) }

func (it *memSwap) Collect() (out chan Result) {
	out = make(chan Result, it.Len())
	defer close(out)

	stat, e := psutil.SwapMemory()
	for _, metric := range it.metrics {
		out <- metric.Collect(stat, e)
	}

	return
}

func MemSwap() *memSwap { return new(memSwap) }
