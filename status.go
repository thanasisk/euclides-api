package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
)

func RawMemoryHandler(w http.ResponseWriter, r *http.Request) {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	w.WriteHeader(http.StatusOK)
	payload, err := json.Marshal(stats)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if cfg.Debug == true {
			fmt.Fprintf(w, err.Error())
		} else {
			fmt.Fprintf(w, "500 - Sorry, something went wrong on our side")
		}
	}
	fmt.Fprintf(w, string(payload))

}

type CustomMemStats struct {
	Alloc       uint64 //Alloc
	LiveObjects uint64 // Mallocs - Frees
	Sys         uint64 // Sys
	HeapAlloc   uint64 //HeapAlloc
	HeapObjects uint64 //HeapOBjects
	StackSys    uint64 //StackSys
	MSpanSys    uint64 //MSpanSys
	NumGC       uint32 //NumGC
}

func MemoryHandler(w http.ResponseWriter, r *http.Request) {
	var rawStats runtime.MemStats
	runtime.ReadMemStats(&rawStats)
	var stats CustomMemStats
	stats.Alloc = rawStats.Alloc
	stats.LiveObjects = rawStats.Mallocs - rawStats.Frees
	stats.Sys = rawStats.Sys
	stats.HeapAlloc = rawStats.HeapAlloc
	stats.HeapObjects = rawStats.HeapObjects
	stats.StackSys = rawStats.StackSys
	stats.MSpanSys = rawStats.MSpanSys
	stats.NumGC = rawStats.NumGC
	payload, err := json.Marshal(stats)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if cfg.Debug == true {
			fmt.Fprintf(w, err.Error())
		} else {
			fmt.Fprintf(w, "500 - Sorry, something went wrong on our side")
		}
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(payload))
}

func RoutineHandler(w http.ResponseWriter, r *http.Request) {
	goRoutines := runtime.NumGoroutine()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, strconv.Itoa(goRoutines))
}
