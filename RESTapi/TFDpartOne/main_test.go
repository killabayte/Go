package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

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
		if test.Answer = factorial(test.InputData); test.Answer != test.Expected {
			t.Errorf("Test case %d failed. Input %v! Expected %d, got %d", id, test.InputData, test.Expected, test.Answer)
		}
	}
}

type HttpTestCase struct {
	Name     string
	Numeric  int
	Expected []byte
}

var httpCases = []HttpTestCase{
	{Name: "Zero", Numeric: 0, Expected: []byte("1")},
	{Name: "One", Numeric: 1, Expected: []byte("1")},
	{Name: "Three", Numeric: 3, Expected: []byte("6")},
	{Name: "Five", Numeric: 5, Expected: []byte("120")},
	{Name: "Ten", Numeric: 7, Expected: []byte("5040")},
}

func TestHandleFactorial(t *testing.T) {
	for _, test := range httpCases {
		req := &http.Request{Method: "GET"}
		req.URL, _ = url.Parse(fmt.Sprintf("http://localhost:8080/factorial/%d", test.Numeric))
		w := httptest.NewRecorder()
		handleFactorial(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("Test case %s failed. Expected status code %d, got %d", test.Name, http.StatusOK, w.Code)
		}
		if !bytes.Equal(w.Body.Bytes(), test.Expected) {
			t.Errorf("Test case %s failed. Expected %s, got %s", test.Name, test.Expected, w.Body.Bytes())
		}
	}

}
