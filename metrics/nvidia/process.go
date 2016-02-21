package nvidia

type Process struct {
	Name string `json:"name"`
	PID  uint32 `json:"pid"` //!< Process ID
	//!< Amount of used GPU memory in bytes.
	//! Under WDDM, \ref NVML_VALUE_NOT_AVAILABLE is always reported
	//! because Windows KMD manages all the memory and not the NVIDIA driver
	UsedGPUMemory uint64 `json:"used-gpu-memory"`

	Kind ProcessKind
}
