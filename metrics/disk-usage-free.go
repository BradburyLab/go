package metrics

import psutil "github.com/shirou/gopsutil/disk"

type DiskUsageFreeResult struct {
	v    uint64
	name string
	err  error
}

func (it *DiskUsageFreeResult) set(v uint64) *DiskUsageFreeResult     { it.v = v; return it }
func (it *DiskUsageFreeResult) setName(v string) *DiskUsageFreeResult { it.name = v; return it }
func (it *DiskUsageFreeResult) setErr(v error) *DiskUsageFreeResult   { it.err = v; return it }

func (it *DiskUsageFreeResult) V() interface{} { return it.v }
func (it *DiskUsageFreeResult) Name() string   { return it.name }
func (it *DiskUsageFreeResult) Kind() Kind     { return KIND_DISK_USAGE_FREE }
func (it *DiskUsageFreeResult) OK() bool       { return it.err == nil }
func (it *DiskUsageFreeResult) Err() error     { return it.err }

func NewDiskUsFreeResult() *DiskUsageFreeResult { return new(DiskUsageFreeResult) }

type diskUsageFree struct{}

func (it *diskUsageFree) Collect(stat *psutil.DiskUsageStat, name string, e error) Result {
	result := NewDiskUsFreeResult().
		setName(name)
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.Free)
	}
	return result
}

func DiskUsageFree() *diskUsageFree { return new(diskUsageFree) }
