package cpu

import "strconv"

type HWInfo map[string]string

func (r HWInfo) GetString(key string) string {
	if v, exist := r[key]; exist {
		return v
	}
	return ""
}

func (r HWInfo) GetUint64(key string) uint64 {
	if v, exist := r[key]; exist {
		if i, err := strconv.Atoi(v); err == nil {
			return uint64(i)
		}
	}
	return 0
}
