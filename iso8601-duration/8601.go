package iso8601duration

import (
	"bytes"
	"strconv"
	"time"
)

const (
	P byte = byte('P')
	T byte = byte('T')

	Y byte = byte('Y')
	W byte = byte('W')
	D byte = byte('D')
	H byte = byte('H')
	M byte = byte('M')
	S byte = byte('S')
)

// use case: MPD/DASH playlists
// convert go time.Duration to ISO8601 duration compatible string
func String(v time.Duration) string {
	if v < time.Second {
		return "PT0S"
	}

	y := uint64(v.Hours() / (24 * 365))
	w := uint64((uint64(v.Hours()) - y*24*365) / (24 * 7))
	d := uint64(v.Hours()/24) - w*7 - y*365
	h := uint64(v.Hours()) - d*24 - w*7*24 - y*365*24
	m := uint64(v.Minutes()) - h*60 - d*24*60 - w*7*24*60 - y*365*24*60
	s := uint64(v.Seconds()) - m*60 - h*3600 - d*24*3600 - w*7*24*3600 - y*365*24*3600

	buf := bytes.Buffer{}

	buf.WriteByte(P)
	if y != 0 {
		buf.WriteString(strconv.FormatUint(y, 10))
		buf.WriteByte(Y)
	}
	if w != 0 {
		buf.WriteString(strconv.FormatUint(w, 10))
		buf.WriteByte(W)
	}
	if d != 0 {
		buf.WriteString(strconv.FormatUint(d, 10))
		buf.WriteByte(D)
	}

	if h != 0 || m != 0 || s != 0 {
		buf.WriteByte(T)

		if h != 0 {
			buf.WriteString(strconv.FormatUint(h, 10))
			buf.WriteByte(H)
		}
		if m != 0 {
			buf.WriteString(strconv.FormatUint(m, 10))
			buf.WriteByte(M)
		}
		if s != 0 {
			buf.WriteString(strconv.FormatUint(s, 10))
			buf.WriteByte(S)
		}
	}

	return buf.String()
}
