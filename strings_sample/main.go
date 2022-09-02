package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type RepeatString struct {
	Str string `json:"repeat_str"`
}

// Repeat: 文字列を繰り返して結合
func repeatStrings(w http.ResponseWriter, r *http.Request) {
	str := r.URL.Query().Get("str")
  var mul int
  mulStr := r.URL.Query().Get("num")
	mul, _ = strconv.Atoi(mulStr)
	repStr := RepeatString{strings.Repeat(str, mul)}

  w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repStr)
}

func main() {
	http.HandleFunc("/repeat", repeatStrings)

	log.Println("Listening...")
	http.ListenAndServe("localhost:8000", nil)
}
