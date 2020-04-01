package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"
)

func StackDumpHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(debug.Stack()))
}
func HeapDumpHandler(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	s, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// we are already in a debug context so no need to check
		fmt.Fprintf(w, err.Error())
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(s))
}
func GCStatsHandler(w http.ResponseWriter, r *http.Request) {
	var garC debug.GCStats
	debug.ReadGCStats(&garC)
	payload, err := json.Marshal(garC)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// we are already in a debug context so no need to check
		fmt.Fprintf(w, err.Error())
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(payload))
}
