package metrics

import psutil "github.com/shirou/gopsutil/mem"

type MemSwapFreeResult struct {
	v   uint64
	err error
}

func (it *MemSwapFreeResult) set(v uint64) *MemSwapFreeResult   { it.v = v; return it }
func (it *MemSwapFreeResult) setErr(v error) *MemSwapFreeResult { it.err = v; return it }

func (it *MemSwapFreeResult) V() interface{} { return it.v }
func (it *MemSwapFreeResult) Name() string   { return "" }
func (it *MemSwapFreeResult) Kind() Kind     { return KIND_MEM_SWAP_FREE }
func (it *MemSwapFreeResult) OK() bool       { return it.err == nil }
func (it *MemSwapFreeResult) Err() error     { return it.err }

func NewMemSwapFreeResult() *MemSwapFreeResult { return new(MemSwapFreeResult) }

type memSwapFree struct{}

func (it *memSwapFree) Collect(stat *psutil.SwapMemoryStat, e error) Result {
	result := NewMemSwapFreeResult()
	if e != nil {
		return result.setErr(e)
	} else {
		result.set(stat.Free)
	}
	return result
}

func MemSwapFree() *memSwapFree { return new(memSwapFree) }
