package codec

import "errors"

var (
	ErrUnknownKind       = errors.New("unknown codec kind")
	ErrItNotProtobufable = errors.New("structure is not protobuffable")
)
