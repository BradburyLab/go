package metrics

import (
	"time"

	psutil "github.com/shirou/gopsutil/cpu"
)

type CPUResult struct {
	v float64
	e error
}

func (it *CPUResult) set(v float64) *CPUResult  { it.v = v; return it }
func (it *CPUResult) setErr(v error) *CPUResult { it.e = v; return it }

func (it *CPUResult) V() interface{} { return it.v }
func (it *CPUResult) Name() string   { return "" }
func (it *CPUResult) Kind() Kind     { return KIND_CPU }
func (it *CPUResult) OK() bool       { return true }
func (it *CPUResult) Err() error     { return it.e }

func NewCPUResult() *CPUResult { return new(CPUResult) }

type cpu struct{}

func (it *cpu) Len() int { return 1 }

func (it *cpu) Collect() (out chan Result) {
	out = make(chan Result, 1)
	defer close(out)

	result := NewCPUResult()
	ret, e := psutil.CPUPercent(1*time.Second, false)
	if e != nil {
		result.setErr(e)
	} else {
		result.set(ret[0])
	}

	out <- result
	return
}

func CPU() *cpu { return new(cpu) }
