package isitup

import (
	"testing"
)

func TestGetDomain(t *testing.T) {
	v := GetDomain("example.com", "go-isitup Unit Tests")
	if v != "example.com" {
		t.Error("Expected example.com, got", v)
	}
}

func TestGetIP(t *testing.T) {
	v := GetIP("example.com", "go-isitup Unit Tests")
	if v != "93.184.216.34" {
		t.Error("Expected 93.184.216.34, got", v)
	}
}

func TestGetPort(t *testing.T) {
	v := GetPort("example.com", "go-isitup Unit Tests")
	if v != 80 {
		t.Error("Expected 80, got", v)
	}
}

func TestGetStatusCode(t *testing.T) {
	v := GetStatusCode("example.com", "go-isitup Unit Tests")
	if v != 200 {
		t.Error("Expected HTTP code 200, got", v)
	}
}

func TestGetResponseTime(t *testing.T) {
	//	v := GetResponseTime("example.com", "go-isitup Unit Tests")
	// Empty test for now.
	// Not sure _how_ to test, response times are variable!
}
