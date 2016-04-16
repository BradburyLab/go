package metrics

import psutil "github.com/shirou/gopsutil/disk"

type DiskUsageTotalResult struct {
	v    uint64
	name string
	err  error
}

func (it *DiskUsageTotalResult) set(v uint64) *DiskUsageTotalResult     { it.v = v; return it }
func (it *DiskUsageTotalResult) setName(v string) *DiskUsageTotalResult { it.name = v; return it }
func (it *DiskUsageTotalResult) setErr(v error) *DiskUsageTotalResult   { it.err = v; return it }

func (it *DiskUsageTotalResult) V() interface{} { return it.v }
func (it *DiskUsageTotalResult) Name() string   { return it.name }
func (it *DiskUsageTotalResult) Kind() Kind     { return KIND_DISK_USAGE_TOTAL }
func (it *DiskUsageTotalResult) OK() bool       { return it.err == nil }
func (it *DiskUsageTotalResult) Err() error     { return it.err }

func NewDiskUsageTotalResult() *DiskUsageTotalResult { return new(DiskUsageTotalResult) }

type diskUsageTotal struct{}

func (it *diskUsageTotal) Collect(stat *psutil.UsageStat, name string, e error) Result {
	result := NewDiskUsageTotalResult().
		setName(name)
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.Total)
	}
	return result
}

func DiskUsageTotal() *diskUsageTotal { return new(diskUsageTotal) }
