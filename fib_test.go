package main

import (
	"testing"
)

type testValue struct {
	value  int
	actual int
}

func TestFib(t *testing.T) {
	testValues := []testValue{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{5, 5},
		{10, 55},
		{20, 6765},
	}

	for _, data := range testValues {
		fib, _ := fib(data.value)
		if fib != data.actual {
			t.Errorf("got %v\n want %v", fib, data.actual)
		}
	}
}
