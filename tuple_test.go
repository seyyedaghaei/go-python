package python

import "testing"

func TestTuple(t *testing.T) {
	defer Finalize()
	tuple := NewTuple(1)
	tuple.SetItem(0, 5)
	if tuple.GetItem(0).AsInt().Int() != 5 {
		t.Fail()
	}
}
