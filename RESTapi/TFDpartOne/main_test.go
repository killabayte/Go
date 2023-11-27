package main

import "testing"

type TestCase struct {
	Input    int
	Answer   int
	Expected int
}

var cases []TestCase = []TestCase{
	{Input: 0, Expected: 1},
	{Input: 1, Expected: 1},
	{Input: 3, Expected: 6},
	{Input: 5, Expected: 120},
}

func TestFactorial(t *testing.T) {
	for _, test := range cases {
		test.Answer = Factorial(test.Input)
		if test.Answer != test.Expected {
			t.Errorf("Factorial(%d) = %d, want %d", test.Input, test.Answer, test.Expected)
		}
	}
}
