package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouter(t *testing.T) {
	// routerの作成
	mux := NewRouter(context.Background())
	// 上で作成したrouterで、テスト用サーバーを起動
	ts := httptest.NewServer(mux)
	// テスト用リクエストの作成
	r, err := http.NewRequest(http.MethodGet, ts.URL+"/test/100", nil)
	if err != nil {
		t.Fatal(err)
	}
	want := "testId is 100"

	// リクエストの実行、レスポンスの取得
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスボディの読み取り
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if string(respBody) != want {
		t.Errorf("got %q, want %q", respBody, want)
	}
}
