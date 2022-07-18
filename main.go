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
	log.Print("Application started. Listening for requests.")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func getTime(w http.ResponseWriter, r *http.Request) {
	timezones := r.URL.Query().Get("tz")

	resp := make(map[string]string)
	for _, tz := range strings.Split(timezones, ",") {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Invalid timezone")
			log.Print("Invalid timezone")
			return
		}

		if tz == "" {
			tz = "current_time"
		}
		log.Printf("timezone: %s", tz)
		resp[tz] = time.Now().In(loc).String()
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
