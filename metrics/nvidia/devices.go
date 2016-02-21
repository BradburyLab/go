package nvidia

import (
	"bytes"
	"encoding/json"
	"io"
)

type Devices []*Device

func devicesDecode(data *bytes.Buffer) (Devices, error) {
	it := NewDevices()
	if e := json.Unmarshal(data.Bytes(), &it); e != nil {
		return nil, Errorf(ERROR_3, e.Error())
	}
	return it, nil
}

func DevicesDecodeFrom(reader io.Reader) (Devices, error) {
	data := new(bytes.Buffer)
	if _, e := data.ReadFrom(reader); e != nil {
		return nil, Errorf(ERROR_2, e.Error())
	}

	return devicesDecode(data)
}

func (it Devices) SelectMinMemoryInfoPUsed() *Device {
	var model *Device = nil

	if it == nil || len(it) == 0 {
		return model
	}

	for _, r := range it {
		if model == nil {
			model = r
		} else if r.MemoryInfoPUsed() <= model.MemoryInfoPUsed() {
			model = r
		}
	}

	return model
}

func (it Devices) SelectMaxMemoryInfoPFree() *Device {
	var model *Device = nil

	if it == nil || len(it) == 0 {
		return model
	}

	for _, r := range it {
		if model == nil {
			model = r
		} else if r.MemoryInfoPFree() >= model.MemoryInfoPFree() {
			model = r
		}
	}

	return model
}

func (it Devices) SelectMaxMemoryInfoFree() *Device {
	var model *Device = nil

	if it == nil || len(it) == 0 {
		return model
	}

	for _, r := range it {
		if model == nil {
			model = r
		} else if r.MemoryInfoFree >= model.MemoryInfoFree {
			model = r
		}
	}

	return model
}

func (it Devices) SelectMinMemoryInfoUsed() *Device {
	var model *Device = nil

	if it == nil || len(it) == 0 {
		return model
	}

	for _, r := range it {
		if model == nil {
			model = r
		} else if r.MemoryInfoUsed <= model.MemoryInfoUsed {
			model = r
		}
	}

	return model
}

func (it Devices) EncoderUtilizationAVG() uint32 {
	if len(it) == 0 {
		return 0
	}

	var sum uint32 = 0
	for _, model := range it {
		sum += uint32(model.EncoderUtilization)
	}

	return uint32(sum / uint32(len(it)))
}

func NewDevices() Devices { return make(Devices, 0) }
