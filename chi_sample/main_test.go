package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouter(t *testing.T) {
	mux := NewRouter(context.Background())
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/test/100", nil)
	if err != nil {
		t.Fatal(err)
	}

	mux.ServeHTTP(w, r)
	resp := w.Result()
	if err != nil {
		t.Fatal(err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	want := "testId is 100"

	if string(respBody) != want {
		t.Errorf("got %q, want %q", respBody, want)
	}
}
