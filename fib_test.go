package main

import (
	"testing"
)

type testValue struct {
	value  int
	actual int
}

func TestFib(t *testing.T) {
	var testValues []testValue = []testValue{
		{value: 0, actual: 0},
		{value: 1, actual: 1},
		{value: 2, actual: 1},
		{value: 3, actual: 2},
		{value: 5, actual: 5},
		{value: 10, actual: 55},
		{value: 20, actual: 6765},
	}

	for _, data := range testValues {
		if fib(data.value) != data.actual {
			t.Errorf("got %v\n want %v", fib(data.value), data.actual)
		}
	}
}
