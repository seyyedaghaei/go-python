package python

import (
	"testing"
)

func TestDict(t *testing.T) {
	defer Finalize()
	dict := NewDict()
	dict.SetItem("key", "value")
	if dict.GetItem("key").AsString().String() != "value" {
		t.Fail()
	}
}
