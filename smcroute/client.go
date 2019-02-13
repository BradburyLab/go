package smcroute

import (
	"bytes"
	"io"
	"net"
	"time"
)

type Client struct {
	socketPath string
	conn       net.Conn
}

func (it *Client) SetSocketPath(v string) *Client { it.socketPath = v; return it }

func (it *Client) Conn() (net.Conn, *Message) {
	if it.conn == nil {
		if c, e := net.Dial(DEFAULT_NETWORK, it.socketPath); e != nil {
			return nil, Errorf(ERROR_SOCKET_CONNECT, it.socketPath, e.Error())
		} else {
			it.conn = c
		}
	}

	return it.conn, nil
}

func (it *Client) Exec(cmd *CMD) (*bytes.Buffer, *Message) {
	start := time.Now()

	c, err := it.Conn()
	if err != nil {
		return nil, err
	}

	// TODO: reuse connection
	//       now:
	//       1. Exec => OK
	//       2. Exec => hangs:
	//            ...
	//            write(3, "(\0\0\0\0\0\0\0j\0\2\0\0\0\0\0eth0.33\000239.255."..., 40) = 40
	//            read(3, 0xc208060200, 255)              = -1 EAGAIN (Resource temporarily unavailable)
	//            epoll_wait(4, {}, 128, 0)               = 0
	//            epoll_wait(4,
	defer func() {
		c.Close()
		it.conn = nil
	}()

	data, e := cmd.Encode()
	if e != nil {
		return nil, Errorf(ERROR_CMD_ENCODE, cmd.StringBash(), e.Error())
	}

	if _, e := io.Copy(c, data); e != nil {
		return nil, Errorf(ERROR_SOCKET_WRITE, cmd.StringBash(), it.socketPath, e.Error())
	}

	respBuf := make([]byte, RESPONSE_BUFFER_SIZE)
	readed, e := c.Read(respBuf)
	if e != nil {
		return nil, Errorf(ERROR_SOCKET_READ, cmd.StringBash(), it.socketPath, e.Error())
	}
	respBuf = bytes.Trim(respBuf, NULL_CHARACTER_STRING)
	resp := bytes.NewBuffer(respBuf)

	latency := time.Now().Sub(start)

	log().Infof(`{"cmd": "%s", "response": "%s", "latency": "%s"}`,
		cmd.StringBash(), resp.String(), latency)

	if readed != 1 || resp.Len() != 0 {
		// <match>
		if reErrorDropMembershipFailed99.Match(resp.Bytes()) {
			return resp, Errorf(ERROR_DROP_MEMBERSHIP_FAILED_99)
		} else if reErrorFailedLeaveNotAMember.Match(resp.Bytes()) {
			return resp, Errorf(ERROR_FAILED_LEAVE_NOT_A_MEMBER, resp.String())
		}
		// </match>

		return resp, Errorf(ERROR_EXEC, cmd.StringBash(), latency, resp.String())
	}

	return resp, nil
}

func NewClient() *Client {
	it := new(Client)
	it.socketPath = DEFAULT_SOCKET_PATH
	return it
}
