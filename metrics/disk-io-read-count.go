package metrics

import psutil "github.com/shirou/gopsutil/disk"

type DiskIOReadCountResult struct {
	v    uint64
	err  error
	name string
}

func (it *DiskIOReadCountResult) set(v uint64) *DiskIOReadCountResult     { it.v = v; return it }
func (it *DiskIOReadCountResult) setName(v string) *DiskIOReadCountResult { it.name = v; return it }
func (it *DiskIOReadCountResult) setErr(v error) *DiskIOReadCountResult   { it.err = v; return it }

func (it *DiskIOReadCountResult) V() interface{} { return it.v }
func (it *DiskIOReadCountResult) Name() string   { return it.name }
func (it *DiskIOReadCountResult) Kind() Kind     { return KIND_DISK_IO_READ_COUNT }
func (it *DiskIOReadCountResult) OK() bool       { return it.err == nil }
func (it *DiskIOReadCountResult) Err() error     { return it.err }

func NewDiskIOReadCountResult() *DiskIOReadCountResult { return new(DiskIOReadCountResult) }

type diskIOReadCount struct{}

func (it *diskIOReadCount) Collect(stat *psutil.IOCountersStat, name string, e error) Result {
	result := NewDiskIOReadCountResult()
	result.setName(name)
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.ReadCount)
	}
	return result
}

func DiskIOReadCount() *diskIOReadCount { return new(diskIOReadCount) }
