package labrador

import (
	"strconv"
	"strings"
)

// ByName : implements sort.Interface for []PodData based on podName
type ByName []PodData

// ByCPU : implements sort.Interface for []PodData based on cpu
type ByCPU []PodData

// ByMemory : implements sort.Interface for []PodData based on memory
type ByMemory []PodData

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].podName < a[j].podName }

func (a ByCPU) Len() int      { return len(a) }
func (a ByCPU) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCPU) Less(i, j int) bool {
	iCPUStr := strings.Split(a[i].cpu, "m")[0]
	iCPUInt, _ := strconv.Atoi(iCPUStr)

	jCPUStr := strings.Split(a[j].cpu, "m")[0]
	jCPUInt, _ := strconv.Atoi(jCPUStr)

	return iCPUInt < jCPUInt
}

func (a ByMemory) Len() int      { return len(a) }
func (a ByMemory) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByMemory) Less(i, j int) bool {
	iMemStr := strings.Split(a[i].memory, "Mi")[0]
	iMemInt, _ := strconv.Atoi(iMemStr)

	jMemStr := strings.Split(a[j].memory, "Mi")[0]
	jMemInt, _ := strconv.Atoi(jMemStr)

	return iMemInt < jMemInt
}
