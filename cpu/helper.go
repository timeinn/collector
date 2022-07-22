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

func (r HWInfo) GetFloat64(key string) float64 {
	if v, exist := r[key]; exist {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}
	return 0
}
