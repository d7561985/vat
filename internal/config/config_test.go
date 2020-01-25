package config

import (
	"testing"
)

func TestGet(t *testing.T) {
	v := Get()
	if v == nil{
		t.Errorf("nil config")
	}
}
