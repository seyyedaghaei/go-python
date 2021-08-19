package python

import (
	"testing"
)

func TestInt(t *testing.T) {
	defer Finalize()
	num := 1
	i := PyInt(num)
	if i.Int() != num {
		t.Fail()
	}
}
