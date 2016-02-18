package metrics

import psutil "github.com/shirou/gopsutil/mem"

type MemVirtFreeResult struct {
	v   uint64
	err error
}

func (it *MemVirtFreeResult) set(v uint64) *MemVirtFreeResult   { it.v = v; return it }
func (it *MemVirtFreeResult) setErr(v error) *MemVirtFreeResult { it.err = v; return it }

func (it *MemVirtFreeResult) V() interface{} { return it.v }
func (it *MemVirtFreeResult) Name() string   { return "" }
func (it *MemVirtFreeResult) Kind() Kind     { return KIND_MEM_VIRT_FREE }
func (it *MemVirtFreeResult) OK() bool       { return it.err == nil }
func (it *MemVirtFreeResult) Err() error     { return it.err }

func NewMemVirtFreeResult() *MemVirtFreeResult { return new(MemVirtFreeResult) }

type memVirtFree struct{}

func (it *memVirtFree) Collect(stat *psutil.VirtualMemoryStat, e error) Result {
	result := NewMemVirtFreeResult()
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.Free)
	}
	return result
}

func MemVirtFree() *memVirtFree { return new(memVirtFree) }
