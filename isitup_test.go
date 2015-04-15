package isitup

import (
	"testing"
)

func TestGetDomain(t *testing.T) {
	var v string
	v = GetDomain("bbc.co.uk", "go-isitup Unit Tests")
	if v != "bbc.co.uk" {
		t.Error("Expected bbc.co.uk, got", v)
	}
}

func TestGetIP(t *testing.T) {
	var v string
	v = GetIP("example.com", "go-isitup Unit Tests")
	if v != "93.184.216.34" {
		t.Error("Expected 93.184.216.34, got", v)
	}
}

func TestGetPort(t *testing.T) {
	var v int
	v = GetPort("cnn.com", "go-isitup Unit Tests")
	if v != 80 {
		t.Error("Expected 80, got", v)
	}
}
