package main

import (
	"strings"
	"testing"
)

const (
	data     = "test"
	expected = "var Test []byte = []byte{\n\t0x74, 0x65, 0x73, 0x74, \n}"
)

func TestStaticPack(t *testing.T) {
	b, err := StaticPack(strings.NewReader(data), "test")
	if err != nil {
		t.Errorf("Experienced error: %s\n", err.Error())
	}
	if string(b) != expected {
		t.Errorf("\nExpected: %s\n, recieved: %s\n", expected, string(b))
		t.FailNow()
	}
}
