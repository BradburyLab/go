package metrics

import (
	"fmt"
	"strings"

	psutil "github.com/shirou/gopsutil/disk"
)

type MetricDiskIOResult struct {
	err error
}

func (it *MetricDiskIOResult) setErr(v error) *MetricDiskIOResult { it.err = v; return it }

func (it *MetricDiskIOResult) V() interface{} { return 0 }
func (it *MetricDiskIOResult) Name() string   { return "" }
func (it *MetricDiskIOResult) Kind() Kind     { return KIND_DISK_IO }
func (it *MetricDiskIOResult) OK() bool       { return it.err == nil }
func (it *MetricDiskIOResult) Err() error     { return it.err }

type MetricDisksIO interface {
	Name() string
	Len() int
	Collect(*psutil.DiskIOCountersStat, error) []Result
}

type disksIO struct {
	name    string
	metrics []MetricDisksIO
}

func (it *disksIO) Register(metrics ...MetricDisksIO) *disksIO {
	for _, metric := range metrics {
		it.Append(metric)
	}
	return it
}

func (it *disksIO) Append(metric MetricDisksIO) *disksIO {
	it.metrics = append(it.metrics, metric)
	return it
}

func (it *disksIO) Len() (l int) {
	for _, m := range it.metrics {
		l += m.Len()
	}
	return
}

func (it *disksIO) Collect() (out chan Result) {
	out = make(chan Result, it.Len())
	defer close(out)

	stats, e := psutil.DiskIOCounters()
	for _, metric := range it.metrics {
		if stat, ok := stats[metric.Name()]; ok {
			for _, in := range metric.Collect(&stat, e) {
				out <- in
			}
		} else {
			var disks []string
			for disk, _ := range stats {
				disks = append(disks, disk)
			}
			out <- &MetricDiskIOResult{fmt.Errorf(`missing disk "%s", available are: %s`, metric.Name(), strings.Join(disks, " | "))}
		}
	}

	return
}

func DisksIO() *disksIO { return new(disksIO) }
