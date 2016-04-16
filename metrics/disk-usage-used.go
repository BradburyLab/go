package metrics

import psutil "github.com/shirou/gopsutil/disk"

type DiskUsageUsedResult struct {
	v    uint64
	name string
	err  error
}

func (it *DiskUsageUsedResult) set(v uint64) *DiskUsageUsedResult     { it.v = v; return it }
func (it *DiskUsageUsedResult) setName(v string) *DiskUsageUsedResult { it.name = v; return it }
func (it *DiskUsageUsedResult) setErr(v error) *DiskUsageUsedResult   { it.err = v; return it }

func (it *DiskUsageUsedResult) V() interface{} { return it.v }
func (it *DiskUsageUsedResult) Name() string   { return it.name }
func (it *DiskUsageUsedResult) Kind() Kind     { return KIND_DISK_USAGE_USED }
func (it *DiskUsageUsedResult) OK() bool       { return it.err == nil }
func (it *DiskUsageUsedResult) Err() error     { return it.err }

func NewDiskUsageUsedResult() *DiskUsageUsedResult { return new(DiskUsageUsedResult) }

type diskUsageUsed struct{}

func (it *diskUsageUsed) Collect(stat *psutil.UsageStat, name string, e error) Result {
	result := NewDiskUsageUsedResult().
		setName(name)
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.Used)
	}
	return result
}

func DiskUsageUsed() *diskUsageUsed { return new(diskUsageUsed) }
