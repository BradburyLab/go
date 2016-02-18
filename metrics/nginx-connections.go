package metrics

import libNginx "gl.bradburylab.tv/f451/l/nginx"

type NginxConnectionsResult struct {
	v   uint
	err error
}

func (it *NginxConnectionsResult) set(v uint) *NginxConnectionsResult     { it.v = v; return it }
func (it *NginxConnectionsResult) setErr(v error) *NginxConnectionsResult { it.err = v; return it }

func (it *NginxConnectionsResult) V() interface{} { return it.v }
func (it *NginxConnectionsResult) Name() string   { return "" }
func (it *NginxConnectionsResult) Kind() Kind     { return KIND_NGINX_CONNECTIONS }
func (it *NginxConnectionsResult) OK() bool       { return it.err == nil }
func (it *NginxConnectionsResult) Err() error     { return it.err }

func NewNginxConnectionsResult() *NginxConnectionsResult { return new(NginxConnectionsResult) }

type nginxConnections struct{}

func (it *nginxConnections) Collect(status *libNginx.Status) Result {
	result := NewNginxConnectionsResult()
	if e := status.Err(); e != nil {
		return result.setErr(e)
	}

	return result.set(status.ActiveConnections())
}

func NginxConnections() *nginxConnections { return new(nginxConnections) }
