package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ReturnString struct {
	Str string `json:"str"`
}

type JoinString struct {
	Str1 string
	Str2 string
}

// Repeat: 文字列を繰り返して結合
func repeatStrings(w http.ResponseWriter, r *http.Request) {
	str := r.URL.Query().Get("str")
  var mul int
  mulStr := r.URL.Query().Get("num")
	mul, _ = strconv.Atoi(mulStr)
	repStr := ReturnString{strings.Repeat(str, mul)}

  w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repStr)
}

// ToUpper: 小文字を大文字に変換
func upperStrings(w http.ResponseWriter, r *http.Request) {
	str := r.URL.Query().Get("str")
  res := ReturnString{strings.ToUpper(str)}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Join: 文字列を結合
func joinStrings(w http.ResponseWriter, r *http.Request) {
	str1 := r.URL.Query().Get("str1")
	str2 := r.URL.Query().Get("str2")
	res := ReturnString{strings.Join([]string{str1, str2}, ",")}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Join: 文字列を結合 (jSON形式でstr1, str2をPOST)
// 参考: https://www.twihike.dev/docs/golang-web/json-request
func joinStringsPost(w http.ResponseWriter, r *http.Request) {
	var js JoinString
	if err := json.NewDecoder(r.Body).Decode(&js); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid request body")
		return
	}

	res := ReturnString{strings.Join([]string{js.Str1, js.Str2}, ",")}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/repeat", repeatStrings)
	http.HandleFunc("/upper", upperStrings)
	http.HandleFunc("/join", joinStrings)
	http.HandleFunc("/joinpost", joinStringsPost)

	log.Println("Listening...")
	http.ListenAndServe("localhost:8000", nil)
}
