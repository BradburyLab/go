package metrics

import psutil "github.com/shirou/gopsutil/disk"

type DiskUsageUsedPercentResult struct {
	v    float64
	name string
	err  error
}

func (it *DiskUsageUsedPercentResult) set(v float64) *DiskUsageUsedPercentResult { it.v = v; return it }
func (it *DiskUsageUsedPercentResult) setName(v string) *DiskUsageUsedPercentResult {
	it.name = v
	return it
}
func (it *DiskUsageUsedPercentResult) setErr(v error) *DiskUsageUsedPercentResult {
	it.err = v
	return it
}

func (it *DiskUsageUsedPercentResult) V() interface{} { return it.v }
func (it *DiskUsageUsedPercentResult) Name() string   { return it.name }
func (it *DiskUsageUsedPercentResult) Kind() Kind     { return KIND_DISK_USAGE_USED_PERCENT }
func (it *DiskUsageUsedPercentResult) OK() bool       { return it.err == nil }
func (it *DiskUsageUsedPercentResult) Err() error     { return it.err }

func NewDiskUsageUsedPercentResult() *DiskUsageUsedPercentResult {
	return new(DiskUsageUsedPercentResult)
}

type diskUsageUsedPercent struct{}

func (it *diskUsageUsedPercent) Collect(stat *psutil.UsageStat, name string, e error) Result {
	result := NewDiskUsageUsedPercentResult().
		setName(name)
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.UsedPercent)
	}
	return result
}

func DiskUsageUsedPercent() *diskUsageUsedPercent { return new(diskUsageUsedPercent) }
