package metrics

import psutil "github.com/shirou/gopsutil/mem"

type MemSwapUsedResult struct {
	v   uint64
	err error
}

func (it *MemSwapUsedResult) set(v uint64) *MemSwapUsedResult   { it.v = v; return it }
func (it *MemSwapUsedResult) setErr(v error) *MemSwapUsedResult { it.err = v; return it }

func (it *MemSwapUsedResult) V() interface{} { return it.v }
func (it *MemSwapUsedResult) Name() string   { return "" }
func (it *MemSwapUsedResult) Kind() Kind     { return KIND_MEM_SWAP_USED }
func (it *MemSwapUsedResult) OK() bool       { return it.err == nil }
func (it *MemSwapUsedResult) Err() error     { return it.err }

func NewMemSwapUsedResult() *MemSwapUsedResult { return new(MemSwapUsedResult) }

type memSwapUsed struct{}

func (it *memSwapUsed) Collect(stat *psutil.SwapMemoryStat, e error) Result {
	result := NewMemSwapUsedResult()
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.Used)
	}
	return result
}

func MemSwapUsed() *memSwapUsed { return new(memSwapUsed) }
