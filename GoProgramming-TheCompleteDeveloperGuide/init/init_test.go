package main

import "testing"

func TestIsValidEmail(t *testing.T) {
	data := "email@example.com"
	if !isValidEmail(data) {
		t.Errorf("IsValidEmail(%v)=false, want true", data)
	}
}

func TestIsValidEmailTable(t *testing.T) {
	table := []struct {
		email string
		want  bool
	}{
		{"email@example.com", true},
		{"missing@tld", false},
	}
	for _, data := range table {
		result := isValidEmail(data.email)
		if result != data.want {
			t.Errorf("%v: %t, want: %t", data.email, result, data.want)
		}
	}
}
