/*
*  Copyright 2018 Dom Rodriguez (shymega)
*
*  Licensed under the Apache License, Version 2.0 (the "License");
*  you may not use this file except in compliance with the License.
*  You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*  distributed under the License is distributed on an "AS IS" BASIS,
*  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*  See the License for the specific language governing permissions and
*  limitations under the License.
*/

package isitup

import (
	"reflect"
	"testing"
)

func TestGetDomain(t *testing.T) {
	v := GetDomain("example.com", "go-isitup Unit tests")
	if v != "example.com" {
		t.Error("Expected example.com, got", v)
	}
}

func TestGetIP(t *testing.T) {
	v := GetIP("example.com", "go-isitup Unit tests")
	if v != "93.184.216.34" {
		t.Error("Expected 93.184.216.34, got", v)
	}
}

func TestGetPort(t *testing.T) {
	v := GetPort("example.com", "go-isitup Unit tests")
	if v != 80 {
		t.Error("Expected 80, got", v)
	}
}

func TestGetStatusCode(t *testing.T) {
	v := GetStatusCode("example.com", "go-isitup Unit tests")
	if v != 200 {
		t.Error("Expected HTTP code 200, got", v)
	}
}

func TestGetResponseTime(t *testing.T) {
	v := GetResponseTime("example.com", "go-isitup Unit tests")
	typ := reflect.TypeOf(v)

	if typ.Kind() != reflect.Float64 {
		t.Error("Expected type of `float64`, got", reflect.TypeOf(v))
	}
}
