package metrics

import psutil "github.com/shirou/gopsutil/mem"

type MemVirtUsedResult struct {
	v   uint64
	err error
}

func (it *MemVirtUsedResult) set(v uint64) *MemVirtUsedResult   { it.v = v; return it }
func (it *MemVirtUsedResult) setErr(v error) *MemVirtUsedResult { it.err = v; return it }

func (it *MemVirtUsedResult) V() interface{} { return it.v }
func (it *MemVirtUsedResult) Name() string   { return "" }
func (it *MemVirtUsedResult) Kind() Kind     { return KIND_MEM_VIRT_USED }
func (it *MemVirtUsedResult) OK() bool       { return it.err == nil }
func (it *MemVirtUsedResult) Err() error     { return it.err }

func NewMemVirtUsedResult() *MemVirtUsedResult { return new(MemVirtUsedResult) }

type memVirtUsed struct{}

func (it *memVirtUsed) Collect(stat *psutil.VirtualMemoryStat, e error) Result {
	result := NewMemVirtUsedResult()
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.Used)
	}
	return result
}

func MemVirtUsed() *memVirtUsed { return new(memVirtUsed) }
