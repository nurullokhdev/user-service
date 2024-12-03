package main

import (
	"encoding/json"
	"net/http"
)

func readJSON(w http.ResponseWriter, r *http.Request, d interface{}) error {
	maxSize := 1_048_578

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxSize))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(d)
}

func writeJSON(w http.ResponseWriter, jsonData any, status int) {
	if err := json.NewEncoder(w).Encode(jsonData); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}
