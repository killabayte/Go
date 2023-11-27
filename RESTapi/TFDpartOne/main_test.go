package main

import "testing"

type TestCase struct {
	InputData int
	Answer    int
	Expected  int
}

var cases []TestCase = []TestCase{
	{InputData: 0, Expected: 1},
	{InputData: 1, Expected: 1},
	{InputData: 3, Expected: 6},
	{InputData: 5, Expected: 120},
}

func TestFactorial(t *testing.T) {
	for id, test := range cases {
		if test.Answer := (test.InputData); test.Answer != test.Expected {
			t.Errorf("Test case %d failed. Expected %d, got %d", id, test.Expected, test.Answer)
		}
	}
}
