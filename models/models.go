package models

type SystemInfo struct {
	Timestamp     string  `json:"timestamp"`
	InsertionDate string  `json:"insertion_date"`
	CPU           float64 `json:"cpu"`
	Memory        uint64  `json:"memory"`
	OS            string  `json:"os"`
	Hostname      string  `json:"hostname"`
	NumCPU        int     `json:"num_cpu"`
	TotalMemory   uint64  `json:"total_memory"`
	FreeMemory    uint64  `json:"free_memory"`
	TotalStorage  uint64  `json:"total_storage"`
	FreeStorage   uint64  `json:"free_storage"`
	StorageDevice string  `json:"storage_device"`
}
