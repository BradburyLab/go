package decoder

import (
	"bytes"
	"io"

	"github.com/BradburyLab/go/codec"
)

func Read(r io.Reader, it interface{}, kind codec.Kind) (int64, error) {
	if r == nil {
		return 0, nil
	}

	buf := new(bytes.Buffer)
	n, e := buf.ReadFrom(r)
	if e != nil {
		return n, e
	}

	return n, DoBuf(buf, it, kind)
}
