package metrics

type ResultGPU struct {
	v float64
	e error
}

func (it *ResultGPU) V() interface{} { return it.v }
func (it *ResultGPU) Name() string   { return "" }
func (it *ResultGPU) Kind() Kind     { return KIND_GPU }
func (it *ResultGPU) OK() bool       { return true }
func (it *ResultGPU) Err() error     { return it.e }

type gpu struct{}

func (it *gpu) Len() int { return 1 }

func (it *gpu) Collect() (out chan Result) {
	out = make(chan Result, 1)
	defer close(out)
	out <- &ResultGPU{48, nil}
	return out
}

func GPU() *gpu { return new(gpu) }
