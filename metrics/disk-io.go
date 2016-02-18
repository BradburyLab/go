package metrics

import (
	psutil "github.com/shirou/gopsutil/disk"
)

type MetricDiskIO interface {
	Collect(*psutil.DiskIOCountersStat, string, error) Result
}

type diskIO struct {
	name    string
	metrics []MetricDiskIO
}

func (it *diskIO) Name() string { return it.name }

func (it *diskIO) Register(metrics ...MetricDiskIO) *diskIO {
	for _, metric := range metrics {
		it.Append(metric)
	}
	return it
}

func (it *diskIO) Append(metric MetricDiskIO) *diskIO {
	it.metrics = append(it.metrics, metric)
	return it
}

func (it *diskIO) Len() int { return len(it.metrics) }

func (it *diskIO) Collect(stat *psutil.DiskIOCountersStat, e error) (out []Result) {
	for _, metric := range it.metrics {
		out = append(out, metric.Collect(stat, it.name, e))
	}

	return
}

func DiskIO(name string) *diskIO {
	it := new(diskIO)
	it.name = name
	return it
}
