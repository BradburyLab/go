package smcroute

import "fmt"

type Message struct {
	Code MessageCode `json:"code,omitempty" yaml:"code"`
	Text string      `json:"text,omitempty" yaml:"text"`
}

func (self *Message) GetCode() MessageCode { return self.Code }
func (self *Message) GetText() string      { return self.Text }

func (self *Message) Is(code MessageCode) bool { return self.Code == code }
func (self *Message) Error() string            { return self.String() }
func (self *Message) String() string           { return self.Text }

func (self *Message) SetCode(v MessageCode) *Message { self.Code = v; return self }
func (self *Message) SetText(v string) *Message      { self.Text = v; return self }

func NewMessageInitialize(it *Message) {}

func Errorf(code MessageCode, params ...interface{}) *Message {
	text := MessageCodeText(code)
	if len(params) > 0 {
		text = fmt.Sprintf(MessageCodeText(code), params...)
	}

	it := &Message{
		Code: code,
		Text: text,
	}
	NewMessageInitialize(it)
	return it
}

func NewMessage() *Message {
	it := new(Message)
	NewMessageInitialize(it)
	return it
}
