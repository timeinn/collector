package cpu_test

import (
	"github.com/sendya/pkg/json"
	"github.com/timeinn/collector/cpu"
	"testing"
)

func TestGet(t *testing.T) {
	info, _, err := cpu.Get()
	if err != nil {
		t.Error(err)
	}
	t.Log(json.ToJSONf(info))
}
