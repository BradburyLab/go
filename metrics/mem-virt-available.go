package metrics

import psutil "github.com/shirou/gopsutil/mem"

type MemVirtAvailableResult struct {
	v   uint64
	err error
}

func (it *MemVirtAvailableResult) set(v uint64) *MemVirtAvailableResult   { it.v = v; return it }
func (it *MemVirtAvailableResult) setErr(v error) *MemVirtAvailableResult { it.err = v; return it }

func (it *MemVirtAvailableResult) V() interface{} { return it.v }
func (it *MemVirtAvailableResult) Name() string   { return "" }
func (it *MemVirtAvailableResult) Kind() Kind     { return KIND_MEM_VIRT_AVAILABLE }
func (it *MemVirtAvailableResult) OK() bool       { return it.err == nil }
func (it *MemVirtAvailableResult) Err() error     { return it.err }

func NewMemVirtAvailableResult() *MemVirtAvailableResult { return new(MemVirtAvailableResult) }

type memVirtAvailable struct{}

func (it *memVirtAvailable) Collect(stat *psutil.VirtualMemoryStat, e error) Result {
	result := NewMemVirtAvailableResult()
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.Available)
	}
	return result
}

func MemVirtAvailable() *memVirtAvailable { return new(memVirtAvailable) }
