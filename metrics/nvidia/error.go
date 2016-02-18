package nvidia

import "fmt"

type Code uint8

const (
	ERROR_1 Code = iota + 1
	ERROR_2
	ERROR_3
)

var errorCodeText = map[Code]string{
	ERROR_1: "#NvidiaSMI failed to GET %s: %s",
	ERROR_2: "#NvidiaSMI failed to read response body from nvidia-smi: %s",
	ERROR_3: "#NvidiaSMI failed to parse JSON body: %s",
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
