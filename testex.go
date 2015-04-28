package sessions

import (
	"net/http"
	"testing"
)

// Test for GH-8 for CookieStore
func TestGH8CookieStore(t *testing.T) {
	originalPath := "/"
	store := NewCookieStore()
	store.Options.Path = originalPath
	req, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatal("failed to create request", err)
	}
	store.Options.Path = "/foo"
	if session.Options.Path != originalPath {
		t.Fatalf("bad session path: got %q, want %q", session.Options.Path, originalPath)
	}
}
