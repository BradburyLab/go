package metrics

import (
	"time"

	libNginx "gl.bradburylab.tv/f451/l/nginx"
)

type NginxRPSResult struct {
	v   float32
	err error
}

func (it *NginxRPSResult) set(v float32) *NginxRPSResult  { it.v = v; return it }
func (it *NginxRPSResult) setErr(v error) *NginxRPSResult { it.err = v; return it }

func (it *NginxRPSResult) V() interface{} { return it.v }
func (it *NginxRPSResult) Name() string   { return "" }
func (it *NginxRPSResult) Kind() Kind     { return KIND_NGINX_RPS }
func (it *NginxRPSResult) OK() bool       { return it.err == nil }
func (it *NginxRPSResult) Err() error     { return it.err }

func NewNginxRPSResult() *NginxRPSResult { return new(NginxRPSResult) }

type nginxRPS struct {
	prv   uint64
	cur   uint64
	prvAt time.Time
	curAt time.Time
}

func (it *nginxRPS) Collect(status *libNginx.Status) Result {
	result := NewNginxRPSResult()
	if e := status.Err(); e != nil {
		return result.setErr(e)
	}

	it.prv = it.cur
	it.prvAt = it.curAt
	it.cur = status.Requests()
	it.curAt = time.Now()

	if it.prv == 0 {
		return result
	}

	diff := it.cur - it.prv
	elapsed := it.curAt.Sub(it.prvAt).Seconds()

	return result.set(float32(diff) / float32(elapsed))
}

func NginxRPS() *nginxRPS { return new(nginxRPS) }
