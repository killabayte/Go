package main

import "testing"

type TestCase struct {
	InputData int
	Answer    int
	Expected  int
}

var cases []TestCase = []TestCase{
	{Input: 0, Expected: 1},
	{Input: 1, Expected: 1},
	{Input: 3, Expected: 6},
	{Input: 5, Expected: 120},
}

func TestFactorial(t *testing.T) {
	for id, test := range cases {
		if res := (test.InputData); res != test.Expected {
			t.Errorf("Test case %d failed. Expected %d, got %d", id, test.Expected, res)
		}
	}
}
