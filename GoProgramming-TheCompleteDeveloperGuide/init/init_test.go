package main

import "testing"

func TestIsValidEmail(t *testing.T) {
	data := "email@example.com"
	if !isValidEmail(data) {
		t.Error("IsValidEmail(%v)=false, want true", data)
	}
}
