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

}
