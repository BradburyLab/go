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
