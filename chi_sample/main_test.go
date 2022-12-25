package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouter(t *testing.T) {
	r := NewRouter(context.Background())
	wr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/test/100", nil)
	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(wr, req)
	resp := wr.Result()
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
