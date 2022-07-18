package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/time", getTime)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func getTime(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	tzMap := make(map[string]string)
	for _, tz := range strings.Split(q.Get("tz"), ",") {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Invalid timezone")
			return
		}

		if tz == "" {
			tz = "current_time"
		}
		log.Println(tz)
		tzMap[tz] = time.Now().In(loc).String()
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tzMap)
}
