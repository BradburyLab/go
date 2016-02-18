package nginx

import (
	"net"
	"net/http"
	"net/url"
	"strconv"
)

type Nginx struct {
	scheme     string
	host       string
	port       int
	pathStatus string
}

func (it *Nginx) Scheme() string                { return it.scheme }
func (it *Nginx) Host() string                  { return it.host }
func (it *Nginx) Port() int                     { return it.port }
func (it *Nginx) PathStatus() string            { return it.pathStatus }
func (it *Nginx) SetScheme(v string) *Nginx     { it.scheme = v; return it }
func (it *Nginx) SetPathStatus(v string) *Nginx { it.pathStatus = v; return it }

func (it *Nginx) Addr() string { return net.JoinHostPort(it.host, strconv.Itoa(it.port)) }

func (it *Nginx) Status() *Status {
	u := url.URL{
		Scheme: it.Scheme(),
		Host:   it.Addr(),
		Path:   it.PathStatus(),
	}
	resp, e := http.Get(u.String())
	if e != nil {
		return NewStatus().SetErr(Errorf(ERROR_1, u.String(), e.Error()))
	}
	defer resp.Body.Close()
	return StatusDecodeFrom(resp.Body)
}

func NewNginx(host string, port int) *Nginx {
	it := new(Nginx)
	it.scheme = "http"
	it.host = host
	it.port = port
	return it
}
