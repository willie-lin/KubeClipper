package csconfig

import "testing"

func TestNewDefaultConfig(t *testing.T) {
	x := NewDefaultConfig()
	x.Dump()
}
