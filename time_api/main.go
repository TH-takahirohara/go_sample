package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type CurrentTime struct {
	Time time.Time `json:"current_time"`
}

func currentTime(w http.ResponseWriter, r *http.Request) {
	tz := r.URL.Query().Get("tz")
	if tz == "" {
		tz = "UTC"
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "invalid timezone")
		return
	}

	cur := CurrentTime{time.Now().In(loc)}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cur)
}

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/api/time", currentTime)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func main() {
	Start()
}
