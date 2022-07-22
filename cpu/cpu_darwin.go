package cpu

import (
	"fmt"
	"github.com/sendya/pkg/json"
	"os/exec"
	"strconv"
	"strings"
)

var cpuMap = map[string]string{
	"machdep.cpu.vendor":       "vendor_id",
	"machdep.cpu.brand_string": "model_name",
	"machdep.cpu.family":       "family",
	"machdep.cpu.model":        "model",
	"machdep.cpu.stepping":     "stepping",

	"hw.physicalcpu":  "cpu_cores",
	"hw.logicalcpu":   "cpu_logical_processors",
	"hw.cpufrequency": "mhz",
}

func getCpuInfo() (cpuInfo HWInfo, err error) {

	cpuInfo = make(HWInfo)

	for option, key := range cpuMap {
		fmt.Println("opt", option)
		out, err := exec.Command("sysctl", "-n", option).Output()
		if err == nil {
			cpuInfo[key] = strings.Trim(string(out), "\n")
		}
	}

	if len(cpuInfo["mhz"]) != 0 {
		mhz, err := strconv.Atoi(cpuInfo["mhz"])
		if err == nil {
			cpuInfo["mhz"] = strconv.Itoa(mhz / 1000000)
		}
	}

	fmt.Println(json.ToJSONf(cpuInfo))

	return
}
