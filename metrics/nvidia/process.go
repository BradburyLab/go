package nvidia

type Process struct {
	Name string `json:"name" bson:"name"`
	PID  uint32 `json:"pid" bson:"pid"` //!< Process ID
	//!< Amount of used GPU memory in bytes.
	//! Under WDDM, \ref NVML_VALUE_NOT_AVAILABLE is always reported
	//! because Windows KMD manages all the memory and not the NVIDIA driver
	UsedGPUMemory uint64 `json:"used-gpu-memory" bson:"used-gpu-memory"`

	GPUUtil uint32 `json:"gpu-util" bson:"gpu-util"` //!< SM (3D/Compute) Util Value
	MemUtil uint32 `json:"mem-util" bson:"mem-util"` //!< Frame Buffer Memory Util Value
	ENCUtil uint32 `json:"enc-util" bson:"enc-util"` //!< Encoder Util Value
	DECUtil uint32 `json:"dec-util" bson:"dec-util"` //!< Decoder Util Value

	Kind ProcessKind
}
