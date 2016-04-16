package metrics

import psutil "github.com/shirou/gopsutil/disk"

type DiskIOReadBytesResult struct {
	v    uint64
	err  error
	name string
}

func (it *DiskIOReadBytesResult) set(v uint64) *DiskIOReadBytesResult     { it.v = v; return it }
func (it *DiskIOReadBytesResult) setName(v string) *DiskIOReadBytesResult { it.name = v; return it }
func (it *DiskIOReadBytesResult) setErr(v error) *DiskIOReadBytesResult   { it.err = v; return it }

func (it *DiskIOReadBytesResult) V() interface{} { return it.v }
func (it *DiskIOReadBytesResult) Name() string   { return it.name }
func (it *DiskIOReadBytesResult) Kind() Kind     { return KIND_DISK_IO_READ_BYTES }
func (it *DiskIOReadBytesResult) OK() bool       { return it.err == nil }
func (it *DiskIOReadBytesResult) Err() error     { return it.err }

func NewDiskIOReadBytesResult() *DiskIOReadBytesResult { return new(DiskIOReadBytesResult) }

type diskIOReadBytes struct{}

func (it *diskIOReadBytes) Collect(stat *psutil.IOCountersStat, name string, e error) Result {
	result := NewDiskIOReadBytesResult()
	result.setName(name)
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.ReadBytes)
	}
	return result
}

func DiskIOReadBytes() *diskIOReadBytes { return new(diskIOReadBytes) }
