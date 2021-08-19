package python

import (
	"testing"
)

func TestFloat(t *testing.T) {
	defer Finalize()
	num := 1.2
	float := PyFloat(num)
	if float.Float64() != num {
		t.Fail()
	}
}
