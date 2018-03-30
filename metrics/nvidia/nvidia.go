package nvidia

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Nvidia struct {
	scheme string
	host   string
	port   int
	path   string

	timeout                   time.Duration
	maxIdleConnectionsPerHost int
	dialTimeout               time.Duration
	dialKeepAlive             time.Duration
	tlsHandshakeTimeout       time.Duration
	tlsInsecureSkipVerify     bool

	c *http.Client
}

func (it *Nvidia) Scheme() string { return it.scheme }
func (it *Nvidia) Host() string   { return it.host }
func (it *Nvidia) Port() int      { return it.port }
func (it *Nvidia) Path() string   { return it.path }

func (it *Nvidia) SetScheme(v string) *Nvidia         { it.scheme = v; return it }
func (it *Nvidia) SetPath(v string) *Nvidia           { it.path = v; return it }
func (it *Nvidia) SetTimeout(v time.Duration) *Nvidia { it.timeout = v; return it }
func (it *Nvidia) SetMaxIdleConnectionsPerHost(v int) *Nvidia {
	it.maxIdleConnectionsPerHost = v
	return it
}
func (it *Nvidia) SetDialTimeout(v time.Duration) *Nvidia   { it.dialTimeout = v; return it }
func (it *Nvidia) SetDialKeepAlive(v time.Duration) *Nvidia { it.dialKeepAlive = v; return it }
func (it *Nvidia) SetTLSHandshakeTimeout(v time.Duration) *Nvidia {
	it.tlsHandshakeTimeout = v
	return it
}
func (it *Nvidia) SetTLSInsecureSkipVerify(v bool) *Nvidia {
	it.tlsInsecureSkipVerify = v
	return it
}

func (it *Nvidia) Addr() string { return net.JoinHostPort(it.host, strconv.Itoa(it.port)) }
func (it *Nvidia) client() *http.Client {
	if it.c == nil {
		transport := &http.Transport{
			MaxIdleConnsPerHost: it.maxIdleConnectionsPerHost,
			Dial: (&net.Dialer{
				Timeout:   it.dialTimeout,
				KeepAlive: it.dialKeepAlive,
			}).Dial,
			TLSHandshakeTimeout: it.tlsHandshakeTimeout,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: it.tlsInsecureSkipVerify,
			},
		}
		it.c = &http.Client{
			Transport: transport,
			Timeout:   it.timeout,
		}
	}

	return it.c
}

func (it *Nvidia) Status() (Devices, error) {
	u := url.URL{
		Scheme: it.Scheme(),
		Host:   it.Addr(),
		Path:   it.Path(),
	}
	resp, e := it.client().Get(u.String())
	if e != nil {
		return nil, Errorf(ERROR_1, u.String(), e.Error())
	}
	defer resp.Body.Close()
	return DevicesDecodeFrom(resp.Body)
}

func NewNvidia(host string, port int) *Nvidia {
	it := new(Nvidia)
	it.scheme = "http"
	it.host = host
	it.port = port
	it.path = "/"

	it.timeout = 1 * time.Second
	it.maxIdleConnectionsPerHost = 16
	it.dialTimeout = 1 * time.Second
	it.dialKeepAlive = 10 * time.Second
	it.tlsHandshakeTimeout = time.Second
	it.tlsInsecureSkipVerify = true

	return it
}
