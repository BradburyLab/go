package metrics

import (
	psutil "github.com/shirou/gopsutil/mem"
)

type MetricMemVirt interface {
	Collect(*psutil.VirtualMemoryStat, error) Result
}

type memVirt struct {
	metrics []MetricMemVirt
}

func (it *memVirt) Register(metrics ...MetricMemVirt) *memVirt {
	for _, metric := range metrics {
		it.Append(metric)
	}
	return it
}

func (it *memVirt) Append(metric MetricMemVirt) *memVirt {
	it.metrics = append(it.metrics, metric)
	return it
}

func (it *memVirt) Len() int { return len(it.metrics) }

func (it *memVirt) Collect() (out chan Result) {
	out = make(chan Result, it.Len())
	defer close(out)

	stat, e := psutil.VirtualMemory()
	for _, metric := range it.metrics {
		out <- metric.Collect(stat, e)
	}

	return
}

func MemVirt() *memVirt { return new(memVirt) }
