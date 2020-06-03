package idutil

import (
	"testing"
	"time"
)

func TestTryCreateNewId(t *testing.T) {

	var r, err = TryCreateNewId()
	if r != 0 || err != nil {
		t.Errorf("result:%v,err:%v,%v,%v", r, err, time.Now().Unix(), time.Now().UnixNano())
	}
}
