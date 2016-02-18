package metrics

import psutil "github.com/shirou/gopsutil/mem"

type MemVirtTotalResult struct {
	v   uint64
	err error
}

func (it *MemVirtTotalResult) set(v uint64) *MemVirtTotalResult   { it.v = v; return it }
func (it *MemVirtTotalResult) setErr(v error) *MemVirtTotalResult { it.err = v; return it }

func (it *MemVirtTotalResult) V() interface{} { return it.v }
func (it *MemVirtTotalResult) Name() string   { return "" }
func (it *MemVirtTotalResult) Kind() Kind     { return KIND_MEM_VIRT_TOTAL }
func (it *MemVirtTotalResult) OK() bool       { return it.err == nil }
func (it *MemVirtTotalResult) Err() error     { return it.err }

func NewMemVirtTotalResult() *MemVirtTotalResult { return new(MemVirtTotalResult) }

type memVirtTotal struct{}

func (it *memVirtTotal) Collect(stat *psutil.VirtualMemoryStat, e error) Result {
	result := NewMemVirtTotalResult()
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.Total)
	}
	return result
}

func MemVirtTotal() *memVirtTotal { return new(memVirtTotal) }
