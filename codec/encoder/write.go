package encoder

import (
	"fmt"
	"io"

	"github.com/BradburyLab/go/codec"
)

func Write(w io.Writer, v interface{}, kind codec.Kind) (int, error) {
	b, e := Do(v, kind)
	if e != nil {
		return 0, fmt.Errorf("Marshal %s error: %s", kind, e.Error())
	}
	return w.Write(b.Bytes())
}
