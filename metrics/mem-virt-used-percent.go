package metrics

import psutil "github.com/shirou/gopsutil/mem"

type MemVirtUsedPercentResult struct {
	v   float64
	err error
}

func (it *MemVirtUsedPercentResult) set(v float64) *MemVirtUsedPercentResult  { it.v = v; return it }
func (it *MemVirtUsedPercentResult) setErr(v error) *MemVirtUsedPercentResult { it.err = v; return it }

func (it *MemVirtUsedPercentResult) V() interface{} { return it.v }
func (it *MemVirtUsedPercentResult) Name() string   { return "" }
func (it *MemVirtUsedPercentResult) Kind() Kind     { return KIND_MEM_VIRT_USED_PERCENT }
func (it *MemVirtUsedPercentResult) OK() bool       { return it.err == nil }
func (it *MemVirtUsedPercentResult) Err() error     { return it.err }

func NewMemVirtUsedPercentResult() *MemVirtUsedPercentResult { return new(MemVirtUsedPercentResult) }

type memVirtUsedPercent struct{}

func (it *memVirtUsedPercent) Collect(stat *psutil.VirtualMemoryStat, e error) Result {
	result := NewMemVirtUsedPercentResult()
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.UsedPercent)
	}
	return result
}

func MemVirtUsedPercent() *memVirtUsedPercent { return new(memVirtUsedPercent) }
