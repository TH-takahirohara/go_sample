package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type ReturnRandom struct {
	Rand int `json:"rand"`
}

// 0~maxnumの間の乱数値(整数)を返す
func randomInt(w http.ResponseWriter, r *http.Request) {
	numStr := r.URL.Query().Get("maxnum")
	num, _ := strconv.Atoi(numStr)
	rand.Seed(time.Now().UnixNano())
  res := ReturnRandom{rand.Intn(num)}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/random", randomInt)

	log.Println("Listening...")
	http.ListenAndServe("localhost:8000", nil)
}
