package nvidia

type Device struct {
	Index              uint8  `json:"index"`
	Name               string `json:"name"`
	MemoryInfoTotal    uint64 `json:"memory-info-total"`
	MemoryInfoFree     uint64 `json:"memory-info-free"`
	MemoryInfoUsed     uint64 `json:"memory-info-used"`
	DecoderUtilization uint8  `json:"decoder-utilization"`
	EncoderUtilization uint8  `json:"encoder-utilization"`
}

func (it *Device) MemoryInfoPFree() float64 {
	if it.MemoryInfoTotal == 0 {
		return 0
	}

	return (float64(it.MemoryInfoFree) / float64(it.MemoryInfoTotal)) * 100
}

func (it *Device) MemoryInfoPUsed() float64 {
	if it.MemoryInfoTotal == 0 {
		return 0
	}

	return (float64(it.MemoryInfoUsed) / float64(it.MemoryInfoTotal)) * 100
}
