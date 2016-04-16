package metrics

import (
	"fmt"

	psutil "github.com/shirou/gopsutil/disk"
)

type MetricDiskUsage interface {
	Collect(*psutil.UsageStat, string, error) Result
}

type diskUsage struct {
	name    string
	metrics []MetricDiskUsage
}

func (it *diskUsage) Name() string { return it.name }

func (it *diskUsage) Register(metrics ...MetricDiskUsage) *diskUsage {
	for _, metric := range metrics {
		it.Append(metric)
	}
	return it
}

func (it *diskUsage) Append(metric MetricDiskUsage) *diskUsage {
	it.metrics = append(it.metrics, metric)
	return it
}

func (it *diskUsage) Len() int { return len(it.metrics) }

func (it *diskUsage) Collect() (out chan Result) {
	out = make(chan Result, it.Len())
	defer close(out)

	stat, e := psutil.Usage(it.Name())
	if e != nil {
		e = fmt.Errorf(`retrieve disk-usage status of "%s" failed: %s`, it.Name(), e.Error())
	}
	for _, metric := range it.metrics {
		out <- metric.Collect(stat, it.name, e)
	}

	return
}

func DiskUsage(name string) *diskUsage {
	it := new(diskUsage)
	it.name = name
	return it
}
