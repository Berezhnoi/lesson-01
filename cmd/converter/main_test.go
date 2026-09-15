package main

import "testing"

func TestParseArgs(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		amount, rate, err := parseArgs([]string{"100", "0.91"})
		if err != nil {
			t.Fatalf("parseArgs returned unexpected error: %v", err)
		}
		if amount != 100 {
			t.Fatalf("amount = %v, want 100", amount)
		}
		if rate != 0.91 {
			t.Fatalf("rate = %v, want 0.91", rate)
		}
	})

	t.Run("invalid amount", func(t *testing.T) {
		_, _, err := parseArgs([]string{"asa", "0.91"})
		if err == nil {
			t.Fatal("parseArgs should return an error for invalid amount")
		}
		if got := err.Error(); got == "" || !contains(got, "invalid amount") || !contains(got, "example: converter") {
			t.Fatalf("unexpected error message: %q", got)
		}
	})

	t.Run("invalid rate", func(t *testing.T) {
		_, _, err := parseArgs([]string{"100", "asa"})
		if err == nil {
			t.Fatal("parseArgs should return an error for invalid rate")
		}
		if got := err.Error(); got == "" || !contains(got, "invalid rate") || !contains(got, "example: converter") {
			t.Fatalf("unexpected error message: %q", got)
		}
	})

	t.Run("missing arguments", func(t *testing.T) {
		_, _, err := parseArgs([]string{"100"})
		if err == nil {
			t.Fatal("parseArgs should return an error when arguments are missing")
		}
		if got := err.Error(); got == "" || !contains(got, "usage: converter") || !contains(got, "example: converter") {
			t.Fatalf("unexpected error message: %q", got)
		}
	})
}

func contains(s, substr string) bool {
	return len(substr) == 0 || (len(s) >= len(substr) && (func() bool {
		for i := 0; i+len(substr) <= len(s); i++ {
			if s[i:i+len(substr)] == substr {
				return true
			}
		}
		return false
	})())
}
