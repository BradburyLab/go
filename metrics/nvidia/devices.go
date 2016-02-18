package nvidia

import (
	"bytes"
	"encoding/json"
	"io"
)

type Devices []*Device

func devicesDecode(data *bytes.Buffer) (Devices, error) {
	it := NewDevices()
	if e := json.Unmarshal(buf.Bytes(), &it); e != nil {
		return Errorf(ERROR_3, e.Error())
	}
	return it, nil
}

func DevicesDecodeFrom(reader io.Reader) (Devices, error) {
	data := new(bytes.Buffer)
	if _, e := data.ReadFrom(reader); e != nil {
		return Errorf(ERROR_2, e.Error())
	}

	return statusDecode(data)
}

func NewDevices() Devices { return make(Devices, 0) }
