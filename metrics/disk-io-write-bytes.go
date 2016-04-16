package metrics

import psutil "github.com/shirou/gopsutil/disk"

type DiskIOWriteBytesResult struct {
	v    uint64
	err  error
	name string
}

func (it *DiskIOWriteBytesResult) set(v uint64) *DiskIOWriteBytesResult     { it.v = v; return it }
func (it *DiskIOWriteBytesResult) setName(v string) *DiskIOWriteBytesResult { it.name = v; return it }
func (it *DiskIOWriteBytesResult) setErr(v error) *DiskIOWriteBytesResult   { it.err = v; return it }

func (it *DiskIOWriteBytesResult) V() interface{} { return it.v }
func (it *DiskIOWriteBytesResult) Name() string   { return it.name }
func (it *DiskIOWriteBytesResult) Kind() Kind     { return KIND_DISK_IO_WRITE_BYTES }
func (it *DiskIOWriteBytesResult) OK() bool       { return it.err == nil }
func (it *DiskIOWriteBytesResult) Err() error     { return it.err }

func NewDiskIOWriteBytesResult() *DiskIOWriteBytesResult { return new(DiskIOWriteBytesResult) }

type diskIOWriteBytes struct{}

func (it *diskIOWriteBytes) Collect(stat *psutil.IOCountersStat, name string, e error) Result {
	result := NewDiskIOWriteBytesResult()
	result.setName(name)
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.WriteBytes)
	}
	return result
}

func DiskIOWriteBytes() *diskIOWriteBytes { return new(diskIOWriteBytes) }
