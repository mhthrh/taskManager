package common

import "testing"

func TestConvertPassword(t *testing.T) {
	pass := "123456"
	p := ConvertPassword(pass)
	if p == pass {
		t.Error("convert password error")
	}
	if p == "" {
		t.Error("convert password error")
	}

}
