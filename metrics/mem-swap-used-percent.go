package metrics

import psutil "github.com/shirou/gopsutil/mem"

type MemSwapUsedPercentResult struct {
	v   float64
	err error
}

func (it *MemSwapUsedPercentResult) set(v float64) *MemSwapUsedPercentResult  { it.v = v; return it }
func (it *MemSwapUsedPercentResult) setErr(v error) *MemSwapUsedPercentResult { it.err = v; return it }

func (it *MemSwapUsedPercentResult) V() interface{} { return it.v }
func (it *MemSwapUsedPercentResult) Name() string   { return "" }
func (it *MemSwapUsedPercentResult) Kind() Kind     { return KIND_MEM_SWAP_USED_PERCENT }
func (it *MemSwapUsedPercentResult) OK() bool       { return it.err == nil }
func (it *MemSwapUsedPercentResult) Err() error     { return it.err }

func NewMemSwapUsedPercentResult() *MemSwapUsedPercentResult { return new(MemSwapUsedPercentResult) }

type memSwapUsedPercent struct{}

func (it *memSwapUsedPercent) Collect(stat *psutil.SwapMemoryStat, e error) Result {
	result := NewMemSwapUsedPercentResult()
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.UsedPercent)
	}
	return result
}

func MemSwapUsedPercent() *memSwapUsedPercent { return new(memSwapUsedPercent) }
