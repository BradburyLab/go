package encoder

import (
	"bytes"
	"encoding/json"
	"encoding/xml"

	"github.com/golang/protobuf/proto"
	"gopkg.in/yaml.v2"

	"github.com/BradburyLab/go/codec"
)

func Do(it interface{}, kind codec.Kind) (*bytes.Buffer, error) {
	switch kind {
	case codec.JSON:
		return JSON(it)
	case codec.XML:
		return XML(it)
	case codec.YAML:
		return YAML(it)
	case codec.PROTO:
		if v, ok := it.(proto.Message); !ok {
			return nil, codec.ErrItNotProtobufable
		} else {
			return PROTO(v)
		}
	}

	return new(bytes.Buffer), nil
}

func Must(v interface{}, ext codec.Kind) *bytes.Buffer {
	b, _ := Do(v, ext)
	return b
}

func MustJSON(v interface{}) *bytes.Buffer {
	b, _ := JSON(v)
	return b
}

func JSON(v interface{}) (*bytes.Buffer, error) {
	b, e := json.Marshal(v)
	return bytes.NewBuffer(b), e
}

func JSONPretty(v interface{}) (*bytes.Buffer, error) {
	b, e := json.MarshalIndent(v, " ", " ")
	return bytes.NewBuffer(b), e
}

func PROTO(v proto.Message) (*bytes.Buffer, error) {
	b, e := proto.Marshal(v)
	return bytes.NewBuffer(b), e
}

func XML(v interface{}) (*bytes.Buffer, error) {
	b, e := xml.Marshal(v)
	return bytes.NewBuffer(b), e
}

func YAML(v interface{}) (*bytes.Buffer, error) {
	b, e := yaml.Marshal(v)
	return bytes.NewBuffer(b), e
}
