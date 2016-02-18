package metrics

import libNvidia "github.com/BradburyLab/go/metrics/nvidia"

type NvidiaResult struct {
	v libNvidia.Devices
	e error
}

func (it *NvidiaResult) set(v libNvidia.Devices) *NvidiaResult { it.v = v; return it }
func (it *NvidiaResult) setErr(v error) *NvidiaResult          { it.e = v; return it }

func (it *NvidiaResult) V() interface{} { return it.v }
func (it *NvidiaResult) Name() string   { return "" }
func (it *NvidiaResult) Kind() Kind     { return KIND_NVIDIA }
func (it *NvidiaResult) OK() bool       { return true }
func (it *NvidiaResult) Err() error     { return it.e }

func NewNvidiaResult() *NvidiaResult { return new(NvidiaResult) }

type nvidia struct {
	scheme string
	host   string
	port   int
	path   string
}

func (it *nvidia) SetScheme(v string) *nvidia { it.scheme = v; return it }
func (it *nvidia) SetHost(v string) *nvidia   { it.host = v; return it }
func (it *nvidia) SetPort(v int) *nvidia      { it.port = v; return it }
func (it *nvidia) SetPath(v string) *nvidia   { it.path = v; return it }

func (it *nvidia) Len() int { return 1 }

func (it *nvidia) Collect() (out chan Result) {
	out = make(chan Result, it.Len())
	defer close(out)

	result := NewNvidiaResult()
	devices, e := libNvidia.
		NewNvidia(it.host, it.port).
		SetScheme(it.scheme).
		SetPath(it.path).
		Status()
	if e != nil {
		result.setErr(e)
	} else {
		result.set(devices)
	}

	out <- result
	return
}

func Nvidia() *nvidia {
	it := new(nvidia)

	it.scheme = "http"
	it.host = "127.0.0.1"
	it.port = 4459
	it.path = "/"

	return it
}
