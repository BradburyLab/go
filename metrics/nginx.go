package metrics

import (
	libNginx "github.com/BradburyLab/go/metrics/nginx"
)

type MetricNginx interface {
	Collect(status *libNginx.Status) Result
}

type nginx struct {
	metrics    []MetricNginx
	scheme     string
	host       string
	port       int
	pathStatus string
}

func (it *nginx) SetScheme(v string) *nginx     { it.scheme = v; return it }
func (it *nginx) SetHost(v string) *nginx       { it.host = v; return it }
func (it *nginx) SetPort(v int) *nginx          { it.port = v; return it }
func (it *nginx) SetPathStatus(v string) *nginx { it.pathStatus = v; return it }

func (it *nginx) Register(metrics ...MetricNginx) *nginx {
	for _, metric := range metrics {
		it.Append(metric)
	}
	return it
}

func (it *nginx) Append(metric MetricNginx) *nginx {
	it.metrics = append(it.metrics, metric)
	return it
}

func (it *nginx) Len() int { return len(it.metrics) }

func (it *nginx) Collect() (out chan Result) {
	out = make(chan Result, it.Len())
	defer close(out)

	status := libNginx.
		NewNginx(it.host, it.port).
		SetScheme(it.scheme).
		SetPathStatus(it.pathStatus).
		Status()

	for _, metric := range it.metrics {
		out <- metric.Collect(status)
	}

	return
}

func Nginx() *nginx {
	it := new(nginx)

	it.scheme = "http"
	it.host = "127.0.0.1"
	it.port = 88
	it.pathStatus = "/stub-status"

	return it
}
