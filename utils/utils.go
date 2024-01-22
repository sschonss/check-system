package utils

import (
	"os"
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// GetCPUUsage returns the CPU usage percentage.
func GetCPUUsage() (float64, error) {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	return percentages[0], nil
}

// GetMemoryInfo returns the memory information.
func GetMemoryInfo() (*mem.VirtualMemoryStat, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return memInfo, nil
}

// GetDiskInfo returns the disk information.
func GetDiskInfo() (*disk.UsageStat, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	diskInfo, err := disk.Usage(wd)
	if err != nil {
		return nil, err
	}
	return diskInfo, nil
}

// GetOS returns the operating system name.
func GetOS() string {
	return runtime.GOOS
}

// GetHostname returns the hostname of the machine.
func GetHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return hostname, nil
}
