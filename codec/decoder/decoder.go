package decoder

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/golang/protobuf/proto"
	"gopkg.in/yaml.v2"

	"github.com/BradburyLab/go/codec"
)

func DoBuf(buf *bytes.Buffer, it interface{}, kind codec.Kind) error {
	return Do(buf.Bytes(), it, kind)
}

func Do(buf []byte, it interface{}, kind codec.Kind) error {
	switch kind {
	case codec.JSON:
		return JSON(buf, it)
	case codec.XML:
		return XML(buf, it)
	case codec.PROTO:
		return PROTO(buf, it)
	case codec.YAML:
		return YAML(buf, it)
	}

	return fmt.Errorf("failed to unmarshal %s: unsupported extension", kind)
}

func JSON(buf []byte, it interface{}) error { return json.Unmarshal(buf, it) }
func YAML(buf []byte, it interface{}) error { return yaml.Unmarshal(buf, it) }
func XML(buf []byte, it interface{}) error  { return xml.Unmarshal(buf, it) }

func PROTO(buf []byte, it interface{}) error {
	v, ok := it.(proto.Message)
	if !ok {
		return codec.ErrItNotProtobufable
	}
	return proto.Unmarshal(buf, v)
}
