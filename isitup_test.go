package isitup

import (
	"testing"
)

func TestGetDomain(t *testing.T) {
	v := GetDomain("google.com", "go-isitup Unit Tests")
	if v != "google.com" {
		t.Error("Expected google.com, got", v)
	}
}

func TestGetIP(t *testing.T) {
	v := GetIP("google.com", "go-isitup Unit Tests")
	if v != "216.58.213.110" {
		t.Error("Expected 216.58.213.110, got", v)
	}
}

func TestGetPort(t *testing.T) {
	v := GetPort("google.com", "go-isitup Unit Tests")
	if v != 80 {
		t.Error("Expected 80, got", v)
	}
}

func TestGetStatusCode(t *testing.T) {
	v := GetStatusCode("google.com", "go-isitup Unit Tests")
	if v != 302 {
		t.Error("Expected HTTP code 302, got", v)
	}
}

func TestGetResponseTime(t *testing.T) {
	//	v := GetResponseTime("google.com", "go-isitup Unit Tests")
	// Empty test for now.
	// Not sure _how_ to test, response times are variable!
}
