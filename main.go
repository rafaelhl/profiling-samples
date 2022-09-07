package main

import (
	"github.com/goccy/go-json"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"

	"github.com/rafaelhl/profiling-samples/samples"
)

func main() {
	http.HandleFunc("/json/unmarshall", handleJson)
	http.HandleFunc("/timezone", handleTimezone)
	http.HandleFunc("/aloc", handleAloc)

	log.Fatal(http.ListenAndServe(":7777", nil))
}

func timesParam(r *http.Request, w http.ResponseWriter) (int, bool) {
	timesParam := r.URL.Query().Get("times")
	times, err := strconv.Atoi(timesParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unexpected times value"))
		return 0, false
	}
	return times, true
}

func handleJson(w http.ResponseWriter, r *http.Request) {
	times, ok := timesParam(r, w)
	if !ok {
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
}

func handleTimezone(w http.ResponseWriter, r *http.Request) {
	times, ok := timesParam(r, w)
	if !ok {
		return
	}

	result := samples.Timezone(times)
	w.Write([]byte(result.String()))
}

func handleAloc(w http.ResponseWriter, r *http.Request) {
	times, ok := timesParam(r, w)
	if !ok {
		return
	}

	result := samples.Aloc("test", times)
	body, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unexpected times value"))
		return
	}

	w.Write(body)
}
