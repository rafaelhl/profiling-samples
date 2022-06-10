package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/pprof"
	"strconv"

	"github.com/rafaelhl/profiling-samples/samples"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/profile", pprof.Profile)
	mux.HandleFunc("/json/unmarshall", func(w http.ResponseWriter, r *http.Request) {
		timesParam := r.URL.Query().Get("times")
		times, err := strconv.Atoi(timesParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("unexpected times value"))
			return
		}

		result := samples.Unmarshall[map[string]any](times)
		body, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("unexpected error"))
			return
		}

		w.Write(body)
	})

	log.Fatal(http.ListenAndServe(":7777", mux))
}
