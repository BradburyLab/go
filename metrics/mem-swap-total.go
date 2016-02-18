package metrics

import psutil "github.com/shirou/gopsutil/mem"

type MemSwapTotalResult struct {
	v   uint64
	err error
}

func (it *MemSwapTotalResult) set(v uint64) *MemSwapTotalResult   { it.v = v; return it }
func (it *MemSwapTotalResult) setErr(v error) *MemSwapTotalResult { it.err = v; return it }

func (it *MemSwapTotalResult) V() interface{} { return it.v }
func (it *MemSwapTotalResult) Name() string   { return "" }
func (it *MemSwapTotalResult) Kind() Kind     { return KIND_MEM_SWAP_TOTAL }
func (it *MemSwapTotalResult) OK() bool       { return it.err == nil }
func (it *MemSwapTotalResult) Err() error     { return it.err }

func NewMemSwapTotalResult() *MemSwapTotalResult { return new(MemSwapTotalResult) }

type memSwapTotal struct{}

func (it *memSwapTotal) Collect(stat *psutil.SwapMemoryStat, e error) Result {
	result := NewMemSwapTotalResult()
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.Total)
	}
	return result
}

func MemSwapTotal() *memSwapTotal { return new(memSwapTotal) }
