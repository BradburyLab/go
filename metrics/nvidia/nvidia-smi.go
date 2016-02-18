package nvidia

import (
	"net"
	"net/http"
	"net/url"
	"strconv"
)

type Nvidia struct {
	scheme string
	host   string
	port   int
	path   string

	client *http.Client
}

func (it *Nvidia) Scheme() string { return it.scheme }
func (it *Nvidia) Host() string   { return it.host }
func (it *Nvidia) Port() int      { return it.port }
func (it *Nvidia) Path() string   { return it.path }

func (it *Nvidia) SetScheme(v string) *Nvidia { it.scheme = v; return it }
func (it *Nvidia) SetPath(v string) *Nvidia   { it.pathStatus = v; return it }

func (it *Nvidia) Addr() string { return net.JoinHostPort(it.host, strconv.Itoa(it.port)) }

func (it *Nvidia) Status() (Devices, error) {
	u := url.URL{
		Scheme: it.Scheme(),
		Host:   it.Addr(),
		Path:   it.Path(),
	}
	resp, e := it.client.Get(u.String())
	if e != nil {
		return nil, Errorf(ERROR_1, u.String(), e.Error())
	}
	defer resp.Body.Close()
	return DevicesDecodeFrom(resp.Body)
}

func NewNvidia(scheme, host, port, path string) *Nvidia {
	it := new(Nvidia)
	it.scheme = "http"
	it.host = host
	it.port = port

	it.client = new(http.Client)

	return it
}
