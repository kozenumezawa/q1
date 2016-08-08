package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	actual := fib(10)
	expected := 55
	if actual != expected {
		t.Errorf("got %v\n want %v", actual, expected)
	}
}
