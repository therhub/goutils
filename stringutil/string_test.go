package stringutil

import "testing"

func TestNewID(t *testing.T) {

	var r = NewID()
	if r == 0 {
		t.Errorf("result:%v", r)
	}
}
