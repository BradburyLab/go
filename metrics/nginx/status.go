package nginx

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

const STATUS_LINES = 4

var (
	statusReActiveConnections      = regexp.MustCompile(`Active connections: (\d+)`)
	statusReAcceptsHandledRequests = regexp.MustCompile(`(\d+) (\d+) (\d+)`)
	statusReReadingWritingWaiting  = regexp.MustCompile(`Reading: (\d+) Writing: (\d+) Waiting: (\d+)`)
)

// http://nginx.org/ru/docs/http/ngx_http_stub_status_module.html
type Status struct {
	// [en]
	// The current number of active client connections including Waiting connections.
	// [ru]
	// Текущее число активных клиентских соединений, включая Waiting-соединения.
	activeConnections uint

	// [en]
	// The total number of accepted client connections.
	// [ru]
	// Суммарное число принятых клиентских соединений.
	accepts uint64
	// [en]
	// The total number of handled connections.
	// Generally, the parameter value is the same as accepts unless some
	// resource limits have been reached
	// (for example, the worker_connections limit).
	// [ru]
	// Суммарное число обработанных соединений.
	// Как правило, значение этого параметра такое же, как accepts,
	// если не достигнуто какое-нибудь системное ограничение
	// (например, лимит worker_connections).
	handled uint64
	// [en]
	// The total number of client requests.
	// [ru]
	// Суммарное число клиентских запросов.
	requests uint64

	// [en]
	// The current number of connections where
	// nginx is reading the request header.
	// [ru]
	// Текущее число соединений, в которых
	// nginx в настоящий момент читает заголовок запроса.
	reading uint
	// [en]
	// The current number of connections where
	// nginx is writing the response back to the client.
	// [ru]
	// Текущее число соединений, в которых nginx
	// в настоящий момент отвечает клиенту.
	writing uint
	// [en]
	// The current number of idle client
	// connections waiting for a request.
	// [ru]
	// Текущее число бездействующих клиентских
	// соединений в ожидании запроса.
	waiting uint

	err *Error
}

func (it *Status) ActiveConnections() uint { return it.activeConnections }
func (it *Status) Accepts() uint64         { return it.accepts }
func (it *Status) Handled() uint64         { return it.handled }
func (it *Status) Requests() uint64        { return it.requests }
func (it *Status) Reading() uint           { return it.reading }
func (it *Status) Writing() uint           { return it.writing }
func (it *Status) Waiting() uint           { return it.waiting }

func (it *Status) SetErr(v *Error) *Status { it.err = v; return it }
func (it *Status) OK() bool                { return it.err == nil }
func (it *Status) Err() *Error             { return it.err }

func (it *Status) String() string {
	return fmt.Sprintf(
		`{"active-connections": %d, "accepts": %d, "handled": %d, "requests": %d, "reading": %d, "writing": %d, "waiting": %d}`,
		it.activeConnections,
		it.accepts, it.handled, it.requests,
		it.reading, it.writing, it.waiting,
	)
}

func StatusDecodeFrom(reader io.Reader) *Status {
	data := new(bytes.Buffer)
	if _, e := data.ReadFrom(reader); e != nil {
		return NewStatus().SetErr(Errorf(ERROR_2, e.Error()))
	}

	return statusDecode(data)
}

func statusDecode(data *bytes.Buffer) *Status {
	it := NewStatus()
	scanner := bufio.NewScanner(data)
	line := 0
	for scanner.Scan() {
		if e := scanner.Err(); e != nil {
			return it.SetErr(Errorf(ERROR_13, e.Error()))
		}

		line++

		switch line {
		case 1:
			if subs := statusReActiveConnections.FindStringSubmatch(scanner.Text()); len(subs) != 0 {
				if v, e := strconv.ParseInt(subs[1], 10, 32); e != nil {
					return it.SetErr(Errorf(ERROR_3, subs[1], e.Error()))
				} else {
					it.activeConnections = uint(v)
				}
			} else {
				return it.SetErr(Errorf(ERROR_4, scanner.Text()))
			}

		case 3:
			if subs := statusReAcceptsHandledRequests.FindStringSubmatch(scanner.Text()); len(subs) != 0 {
				if v, e := strconv.ParseInt(subs[1], 10, 64); e != nil {
					return it.SetErr(Errorf(ERROR_5, subs[1], e.Error()))
				} else {
					it.accepts = uint64(v)
				}

				if v, e := strconv.ParseInt(subs[2], 10, 64); e != nil {
					return it.SetErr(Errorf(ERROR_6, subs[2], e.Error()))
				} else {
					it.handled = uint64(v)
				}

				if v, e := strconv.ParseInt(subs[3], 10, 64); e != nil {
					return it.SetErr(Errorf(ERROR_7, subs[3], e.Error()))
				} else {
					it.requests = uint64(v)
				}
			} else {
				Errorf(ERROR_8, scanner.Text())
			}

		case 4:
			if subs := statusReReadingWritingWaiting.FindStringSubmatch(scanner.Text()); len(subs) != 0 {
				if v, e := strconv.ParseInt(subs[1], 10, 32); e != nil {
					return it.SetErr(Errorf(ERROR_9, subs[1], e.Error()))
				} else {
					it.reading = uint(v)
				}

				if v, e := strconv.ParseInt(subs[2], 10, 32); e != nil {
					return it.SetErr(Errorf(ERROR_10, subs[1], e.Error()))
				} else {
					it.writing = uint(v)
				}

				if v, e := strconv.ParseInt(subs[3], 10, 32); e != nil {
					return it.SetErr(Errorf(ERROR_11, subs[1], e.Error()))
				} else {
					it.waiting = uint(v)
				}
			} else {
				return it.SetErr(Errorf(ERROR_12, scanner.Text()))
			}
		}
	}

	if line != STATUS_LINES {
		return it.SetErr(Errorf(ERROR_14, STATUS_LINES, line))
	}

	return it
}

func NewStatus() *Status { return new(Status) }
