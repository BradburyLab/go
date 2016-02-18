package metrics

type Kind uint8

const (
	KIND_CPU Kind = iota
	KIND_NVIDIA
	KIND_LOAD_AVG
	KIND_LOAD_AVG_1
	KIND_LOAD_AVG_5
	KIND_LOAD_AVG_15
	KIND_MEM
	KIND_MEM_VIRT
	KIND_MEM_VIRT_TOTAL
	KIND_MEM_VIRT_AVAILABLE
	KIND_MEM_VIRT_USED
	KIND_MEM_VIRT_USED_PERCENT
	KIND_MEM_VIRT_FREE
	KIND_MEM_SWAP
	KIND_MEM_SWAP_TOTAL
	KIND_MEM_SWAP_FREE
	KIND_MEM_SWAP_USED
	KIND_MEM_SWAP_USED_PERCENT
	KIND_DISK
	KIND_DISK_IO
	KIND_DISK_IO_READ_COUNT
	KIND_DISK_IO_WRITE_COUNT
	KIND_DISK_IO_READ_BYTES
	KIND_DISK_IO_WRITE_BYTES
	KIND_DISK_USAGE
	KIND_DISK_USAGE_TOTAL
	KIND_DISK_USAGE_FREE
	KIND_DISK_USAGE_USED
	KIND_DISK_USAGE_USED_PERCENT
	KIND_NGINX
	KIND_NGINX_CONNECTIONS
	KIND_NGINX_RPS
	KIND_NGINX_LOG
	KIND_NGINX_LOG_VOD
	KIND_NGINX_LOG_VOD_IN
	KIND_NGINX_LOG_VOD_OUT
	KIND_NGINX_LOG_DVR
	KIND_NGINX_LOG_DVR_IN
	KIND_NGINX_LOG_DVR_OUT
)

var kindText = map[Kind]string{
	KIND_CPU:                     "cpu",
	KIND_NVIDIA:                  "nvidia",
	KIND_LOAD_AVG:                "load-avg",
	KIND_LOAD_AVG_1:              "load-avg-1",
	KIND_LOAD_AVG_5:              "load-avg-5",
	KIND_LOAD_AVG_15:             "load-avg-15",
	KIND_MEM:                     "memory",
	KIND_MEM_VIRT:                "memory-virtual",
	KIND_MEM_VIRT_TOTAL:          "memory-virtual-total",
	KIND_MEM_VIRT_AVAILABLE:      "memory-virtual-available",
	KIND_MEM_VIRT_USED:           "memory-virtual-used",
	KIND_MEM_VIRT_USED_PERCENT:   "memory-virtual-used-percent",
	KIND_MEM_VIRT_FREE:           "memory-virtual-free",
	KIND_MEM_SWAP:                "memory-swap",
	KIND_MEM_SWAP_TOTAL:          "memory-swap-total",
	KIND_MEM_SWAP_FREE:           "memory-swap-free",
	KIND_MEM_SWAP_USED:           "memory-swap-used",
	KIND_MEM_SWAP_USED_PERCENT:   "memory-swap-used-percent",
	KIND_DISK_IO:                 "disk-io",
	KIND_DISK_IO_READ_COUNT:      "disk-io-read-count",
	KIND_DISK_IO_WRITE_COUNT:     "disk-io-write-count",
	KIND_DISK_IO_READ_BYTES:      "disk-io-read-bytes",
	KIND_DISK_IO_WRITE_BYTES:     "disk-io-write-bytes",
	KIND_DISK_USAGE:              "disk-usage",
	KIND_DISK_USAGE_TOTAL:        "disk-usage-total",
	KIND_DISK_USAGE_FREE:         "disk-usage-free",
	KIND_DISK_USAGE_USED:         "disk-usage-used",
	KIND_DISK_USAGE_USED_PERCENT: "disk-usage-percent",
	KIND_NGINX:                   "nginx",
	KIND_NGINX_CONNECTIONS:       "nginx-connections",
	KIND_NGINX_RPS:               "nginx-rps",
}

func (it Kind) String() string { return kindText[it] }
