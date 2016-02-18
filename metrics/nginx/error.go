package nginx

import "fmt"

type Code uint8

const (
	ERROR_1  Code = 1
	ERROR_2  Code = 2
	ERROR_3  Code = 3
	ERROR_4  Code = 4
	ERROR_5  Code = 5
	ERROR_6  Code = 6
	ERROR_7  Code = 7
	ERROR_8  Code = 8
	ERROR_9  Code = 9
	ERROR_10 Code = 10
	ERROR_11 Code = 11
	ERROR_12 Code = 12
	ERROR_13 Code = 13
	ERROR_14 Code = 14
)

var errorCodeText = map[Code]string{
	ERROR_1:  "#STATUS failed to GET %s: %s",
	ERROR_2:  "#STATUS failed to read response body after status request: %s",
	ERROR_3:  "#STATUS failed to parse activeConnections from %s: %s",
	ERROR_4:  "#STATUS failed to find activeConnections in %s",
	ERROR_5:  "#STATUS failed to parse accepts from %s: %s",
	ERROR_6:  "#STATUS failed to parse handled from %s: %s",
	ERROR_7:  "#STATUS failed to parse requests from %s: %s",
	ERROR_8:  "#STATUS failed to find accepts, handled, requests in %s",
	ERROR_9:  "#STATUS failed to parse reading from %s: %s",
	ERROR_10: "#STATUS failed to parse writing from %s: %s",
	ERROR_11: "#STATUS failed to parse waiting from %s: %s",
	ERROR_12: "#STATUS failed to find reading, writing, waiting in %s",
	ERROR_13: "#STATUS error while result scanning %s",
	ERROR_14: "#STATUS wrong result line numbers. expected: %d, got: %d",
}

func ErrorCodeText(code Code) string { return errorCodeText[code] }

type Error struct {
	Code Code   `json:"code,omitempty" yaml:"code"`
	Text string `json:"text,omitempty" yaml:"text"`
}

func (self *Error) Error() string  { return fmt.Sprintf("[#%d] %s", self.Code, self.Text) }
func (self *Error) String() string { return self.Text }

func Errorf(code Code, params ...interface{}) *Error {
	return &Error{
		Code: code,
		Text: fmt.Sprintf(ErrorCodeText(code), params...),
	}
}
