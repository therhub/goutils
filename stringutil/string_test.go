package stringutil

import (
	"testing"
)

func TestGetNewId(t *testing.T) {

	var r, err = GetNewId()
	if r == 0 || err != nil {
		t.Errorf("result:%v,err:%v", r, err)
	}
}
