package cpu

import (
	"fmt"
	"strconv"
	"strings"
)

// CPU holds metadata about the host CPU
type CPU struct {
	// VendorId the CPU vendor ID
	VendorId string
	// ModelName the CPU model
	ModelName string
	// CpuCores the number of cores for the CPU
	CpuCores uint64
	// CpuLogicalProcessors the number of logical core for the CPU
	CpuLogicalProcessors uint64
	// Mhz the frequency for the CPU (Not available on ARM)
	Mhz float64
	// CacheSizeBytes the cache size for the CPU (Linux only)
	CacheSizeBytes uint64
	// Family the CPU family
	Family string
	// Model the CPU model name
	Model string
	// Stepping the CPU stepping
	Stepping string

	// CpuPkgs the CPU pkg count (Windows only)
	CpuPkgs uint64
	// CpuNumaNodes the CPU numa node count (Windows only)
	CpuNumaNodes uint64
	// CacheSizeL1Bytes the CPU L1 cache size (Windows only)
	CacheSizeL1Bytes uint64
	// CacheSizeL2Bytes the CPU L2 cache size (Windows only)
	CacheSizeL2Bytes uint64
	// CacheSizeL3 the CPU L3 cache size (Windows only)
	CacheSizeL3Bytes uint64
}

const name = "cpu"

func (r *CPU) Name() string {
	return name
}

// Get returns a Cpu struct already initialized, a list of warnings and an error. The method will try to collect as much
// metadata as possible, an error is returned if nothing could be collected. The list of warnings contains errors if
// some metadata could not be collected.
func Get() (*CPU, []string, error) {
	cpuInfo, err := getCpuInfo()
	if err != nil {
		return nil, nil, err
	}

	var warnings []string
	c := &CPU{}

	c.VendorId = cpuInfo.GetString("vendor_id")
	c.ModelName = cpuInfo.GetString("model_name")
	c.Family = cpuInfo.GetString("family")
	c.Model = cpuInfo.GetString("model")
	c.Stepping = cpuInfo.GetString("stepping")

	// We serialize int to string in the windows version of 'GetCpuInfo' and back to int here. This is less than
	// ideal but we don't want to break backward compatibility for now. The entire gohai project needs a rework but
	// for now we simply adding typed field to avoid using maps of interface..
	c.CpuPkgs = cpuInfo.GetUint64("cpu_pkgs")
	c.CpuNumaNodes = cpuInfo.GetUint64("cpu_numa_nodes")
	c.CacheSizeL1Bytes = cpuInfo.GetUint64("cache_size_l1")
	c.CacheSizeL2Bytes = cpuInfo.GetUint64("cache_size_l2")
	c.CacheSizeL3Bytes = cpuInfo.GetUint64("cache_size_l3")

	c.CpuCores = cpuInfo.GetUint64("cpu_cores")
	c.CpuLogicalProcessors = cpuInfo.GetUint64("cpu_logical_processors")
	c.Mhz = cpuInfo.GetFloat64("mhz")

	// cache_size uses the format '9216 KB'
	cacheSizeString := strings.Split(cpuInfo.GetString("cache_size"), " ")[0]
	cacheSizeBytes, err := strconv.ParseUint(cacheSizeString, 10, 64)
	if err == nil {
		c.CacheSizeBytes = cacheSizeBytes * 1024
	} else {
		warnings = append(warnings, fmt.Sprintf("could not collect cache size: %s", err))
	}

	return c, warnings, nil
}
