package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ReturnString struct {
	Str string `json:"str"`
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

func main() {
	http.HandleFunc("/repeat", repeatStrings)
	http.HandleFunc("/upper", upperStrings)
	http.HandleFunc("/join", joinStrings)

	log.Println("Listening...")
	http.ListenAndServe("localhost:8000", nil)
}
