package login

import "testing"

func TestOidcLogin(t *testing.T) {
	want := "Hello, world."
	if got := OidcLogin(); got != want {
		t.Errorf("OidcLogin() = %q, want = %q", got, want)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want = %q", got, want)
	}
}

