package cpu_test

import (
	"github.com/sendya/pkg/json"
	"github.com/timeinn/collector/cpu"
	"testing"
)

func TestGet(t *testing.T) {
	info, arr, err := cpu.Get()
	if err != nil {
		t.Error(err)
	}

	t.Log(arr)
	t.Log(json.ToJSONf(info))
}
