package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type ReturnSqrt struct {
	Sqrt float64 `json:"sqrt"`
}

type ReturnRandom struct {
	Rand int `json:"rand"`
}

// math.Sqrt: 平方根を返す
func sqrt(w http.ResponseWriter, r *http.Request) {
  numStr := r.URL.Query().Get("num")
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid number")
		return
	}
  res := ReturnSqrt{math.Sqrt(num)}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// 0~maxnumの間の乱数値(整数)を返す
func randomInt(w http.ResponseWriter, r *http.Request) {
	numStr := r.URL.Query().Get("maxnum")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "maxnum must be integer")
		return
	}
	rand.Seed(time.Now().UnixNano())
  res := ReturnRandom{rand.Intn(num)}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/sqrt", sqrt)
	http.HandleFunc("/random", randomInt)

	log.Println("Listening...")
	http.ListenAndServe("localhost:8000", nil)
}
