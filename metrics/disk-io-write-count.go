package metrics

import psutil "github.com/shirou/gopsutil/disk"

type DiskIOWriteCountResult struct {
	v    uint64
	err  error
	name string
}

func (it *DiskIOWriteCountResult) set(v uint64) *DiskIOWriteCountResult     { it.v = v; return it }
func (it *DiskIOWriteCountResult) setName(v string) *DiskIOWriteCountResult { it.name = v; return it }
func (it *DiskIOWriteCountResult) setErr(v error) *DiskIOWriteCountResult   { it.err = v; return it }

func (it *DiskIOWriteCountResult) V() interface{} { return it.v }
func (it *DiskIOWriteCountResult) Name() string   { return it.name }
func (it *DiskIOWriteCountResult) Kind() Kind     { return KIND_DISK_IO_WRITE_COUNT }
func (it *DiskIOWriteCountResult) OK() bool       { return it.err == nil }
func (it *DiskIOWriteCountResult) Err() error     { return it.err }

func NewDiskIOWriteCountResult() *DiskIOWriteCountResult { return new(DiskIOWriteCountResult) }

type diskIOWriteCount struct{}

func (it *diskIOWriteCount) Collect(stat *psutil.IOCountersStat, name string, e error) Result {
	result := NewDiskIOWriteCountResult().
		setName(name)
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.WriteCount)
	}
	return result
}

func DiskIOWriteCount() *diskIOWriteCount { return new(diskIOWriteCount) }
