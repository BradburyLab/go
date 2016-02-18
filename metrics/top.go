package metrics

import (
	"encoding/json"

	libNvidia "github.com/BradburyLab/go/metrics/nvidia"
)

type TopError struct {
	Kind Kind
	Err  error
}

type TopDiskIO struct {
	ReadCount  uint64 `json:"read-count"`
	WriteCount uint64 `json:"write-count"`
	ReadBytes  uint64 `json:"read-bytes"`
	WriteBytes uint64 `json:"write-bytes"`
}

type TopDiskUsage struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used-percent"`
}

type Top struct {
	CPU    float64           `json:"cpu"`
	Nvidia libNvidia.Devices `json:"nvidia"`

	LoadAVG1  float64 `json:"load-avg-1"`
	LoadAVG5  float64 `json:"load-avg-5"`
	LoadAVG15 float64 `json:"load-avg-15"`

	MemVirtTotal       uint64  `json:"mem-virt-total"`
	MemVirtAvailable   uint64  `json:"mem-virt-available"`
	MemVirtUsed        uint64  `json:"mem-virt-used"`
	MemVirtUsedPercent float64 `json:"mem-virt-used-percent"`
	MemVirtFree        uint64  `json:"mem-virt-free"`

	MemSwapTotal       uint64  `json:"mem-swap-total"`
	MemSwapFree        uint64  `json:"mem-swap-free"`
	MemSwapUsed        uint64  `json:"mem-swap-used"`
	MemSwapUsedPercent float64 `json:"mem-swap-used-percent"`

	DiskIO    map[string]*TopDiskIO    `json:"disk-io"`
	DiskUsage map[string]*TopDiskUsage `json:"disk-usage"`

	NginxRPS         float32 `json:"nginx-rps"`
	NginxConnections uint    `json:"nginx-connections"`

	errs []*TopError `json:"-"`
}

func (it *Top) OK() bool          { return it.errs == nil }
func (it *Top) Errs() []*TopError { return it.errs }

func (it *Top) String() string {
	v, _ := json.Marshal(it)
	return string(v)
}

func (it *Top) ensureDiskIO(name string) (out *TopDiskIO) {
	if it.DiskIO == nil {
		it.DiskIO = make(map[string]*TopDiskIO)
	}

	if _, ok := it.DiskIO[name]; !ok {
		it.DiskIO[name] = new(TopDiskIO)
	}

	return it.DiskIO[name]
}

func (it *Top) ensureDiskUsage(name string) (out *TopDiskUsage) {
	if it.DiskUsage == nil {
		it.DiskUsage = make(map[string]*TopDiskUsage)
	}

	if _, ok := it.DiskUsage[name]; !ok {
		it.DiskUsage[name] = new(TopDiskUsage)
	}

	return it.DiskUsage[name]
}

func (it *Top) Reduce(in chan chan Result) *Top {
	for results := range in {
		for result := range results {
			if e := result.Err(); e != nil {
				it.errs = append(it.errs, &TopError{result.Kind(), result.Err()})
			} else if result.OK() {
				switch result.Kind() {
				case KIND_CPU:
					it.CPU = result.V().(float64)
				case KIND_NVIDIA:
					it.Nvidia = result.V().(libNvidia.Devices)

				case KIND_LOAD_AVG_1:
					it.LoadAVG1 = result.V().(float64)
				case KIND_LOAD_AVG_5:
					it.LoadAVG5 = result.V().(float64)
				case KIND_LOAD_AVG_15:
					it.LoadAVG15 = result.V().(float64)

				case KIND_MEM_VIRT_TOTAL:
					it.MemVirtTotal = result.V().(uint64)
				case KIND_MEM_VIRT_AVAILABLE:
					it.MemVirtAvailable = result.V().(uint64)
				case KIND_MEM_VIRT_USED:
					it.MemVirtUsed = result.V().(uint64)
				case KIND_MEM_VIRT_USED_PERCENT:
					it.MemVirtUsedPercent = result.V().(float64)
				case KIND_MEM_VIRT_FREE:
					it.MemVirtFree = result.V().(uint64)

				case KIND_MEM_SWAP_TOTAL:
					it.MemSwapTotal = result.V().(uint64)
				case KIND_MEM_SWAP_FREE:
					it.MemSwapFree = result.V().(uint64)
				case KIND_MEM_SWAP_USED:
					it.MemSwapUsed = result.V().(uint64)
				case KIND_MEM_SWAP_USED_PERCENT:
					it.MemSwapUsedPercent = result.V().(float64)

				case KIND_DISK_IO_READ_COUNT:
					it.ensureDiskIO(result.Name()).ReadCount = result.V().(uint64)
				case KIND_DISK_IO_WRITE_COUNT:
					it.ensureDiskIO(result.Name()).WriteCount = result.V().(uint64)
				case KIND_DISK_IO_READ_BYTES:
					it.ensureDiskIO(result.Name()).ReadBytes = result.V().(uint64)
				case KIND_DISK_IO_WRITE_BYTES:
					it.ensureDiskIO(result.Name()).WriteBytes = result.V().(uint64)

				case KIND_DISK_USAGE_TOTAL:
					it.ensureDiskUsage(result.Name()).Total = result.V().(uint64)
				case KIND_DISK_USAGE_FREE:
					it.ensureDiskUsage(result.Name()).Free = result.V().(uint64)
				case KIND_DISK_USAGE_USED:
					it.ensureDiskUsage(result.Name()).Used = result.V().(uint64)
				case KIND_DISK_USAGE_USED_PERCENT:
					it.ensureDiskUsage(result.Name()).UsedPercent = result.V().(float64)

				case KIND_NGINX_RPS:
					it.NginxRPS = result.V().(float32)
				case KIND_NGINX_CONNECTIONS:
					it.NginxConnections = result.V().(uint)
				}
			}
		}
	}

	return it
}

func NewTop() *Top { return new(Top) }
