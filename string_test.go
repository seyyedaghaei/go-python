package python

import "testing"

func TestString(t *testing.T) {
	defer Finalize()
	if PyString("test").String() != "test" {
		t.Fail()
	}
}
