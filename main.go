package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Debug        bool
	ReadTimeout  time.Duration
	WriteTimeout time.Duration `json:writeTimeout`
	IdleTimeout  time.Duration `json:idleTimeout`
	Port         string
	Address      string
}

var cfg Config

// init is located elsewhere too!
func init() {
	// first and foremost load configuration
	if strings.ToUpper(os.Getenv("DEBUG")) == "TRUE" {
		cfg.Debug = true
	} else {
		cfg.Debug = false
	}
	if len(os.Getenv("ADDRESS")) > 0 {
		cfg.Address = os.Getenv("ADDRESS")
	} else {
		cfg.Address = "0.0.0.0"
	}
	if len(os.Getenv("PORT")) > 0 {
		cfg.Port = os.Getenv("PORT")
	} else {
		cfg.Port = "8080"
	}
	if len(os.Getenv("RDTIMEOUT")) > 0 {
		l, err := strconv.Atoi(os.Getenv("RDTIMEOUT"))
		if err != nil {
			cfg.ReadTimeout = 15
		}
		cfg.ReadTimeout = time.Duration(l)
	} else {
		cfg.ReadTimeout = 15
	}
	if len(os.Getenv("WRTIMEOUT")) > 0 {
		l, err := strconv.Atoi(os.Getenv("WRTIMEOUT"))
		if err != nil {
			cfg.WriteTimeout = 15
		}
		cfg.WriteTimeout = time.Duration(l)
	} else {
		cfg.WriteTimeout = 15
	}
	if len(os.Getenv("IDTIMEOUT")) > 0 {
		l, err := strconv.Atoi(os.Getenv("IDTIMEOUT"))
		if err != nil {
			cfg.IdleTimeout = 60
		}
		cfg.IdleTimeout = time.Duration(l)
	} else {
		cfg.IdleTimeout = 60
	}
}
func main() {
	// no REST API is complete without a nice ASCII logo ...
	logo := `
  ______           _ _     _
 |  ____|         | (_)   | |
 | |__  _   _  ___| |_  __| | ___  ___
 |  __|| | | |/ __| | |/ _  |/ _ \/ __|
 | |___| |_| | (__| | | (_| |  __/\__ \
 |______\__,_|\___|_|_|\__,_|\___||___/
`
	fmt.Print(logo)
	fmt.Println(time.Now().Format(time.RFC850))
	if cfg.Debug == true {
		fmt.Printf("Debug mode - do not run in production")
	} else {
		fmt.Println("Performance mode - certain error messages will be masked")
	}
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/v1/fibonacci/{n:[0-9]+}", NaiveFibonacciHandler)
	r.HandleFunc("/v2/fibonacci/{n:[0-9]+}", SmartFibonacciHandler)
	r.HandleFunc("/v1/ackermann/{n:[0-9]+}/{m:[0-9]+}", NaiveAckermannHandler)
	r.HandleFunc("/v1/factorial/{n:[0-9]+}", NaiveFactorialHandler)
	r.HandleFunc("/v2/factorial/{n:[0-9]+}", SmartFactorialHandler)
	r.HandleFunc("/help", HelpHandler)
	if cfg.Debug == true {
		r.HandleFunc("/status/memory/raw", RawMemoryHandler)
		r.HandleFunc("/debug/stackdump", StackDumpHandler)
		r.HandleFunc("/debug/heapdump", HeapDumpHandler)
		r.HandleFunc("/debug/gcstats", GCStatsHandler)
	}
	r.HandleFunc("/status/memory", MemoryHandler)
	r.HandleFunc("/status/goroutines", RoutineHandler)
	r.Use(loggingMiddleware)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         cfg.Address + ":" + cfg.Port,
		WriteTimeout: cfg.WriteTimeout * time.Second,
		ReadTimeout:  cfg.ReadTimeout * time.Second,
		IdleTimeout:  cfg.IdleTimeout * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

///////////////////////////////////////////////////////////////////////////////
// Fibonacci Section
func NaiveFibonacciHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	candidate, err := strconv.ParseUint(vars["n"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
	}
	res := naiveFibonacci(candidate)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, strconv.FormatUint(res, 10))
}

func SmartFibonacciHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	candidate, err := strconv.ParseUint(vars["n"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "FIXME")
	}
	res := smartFibonacci(candidate)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, strconv.FormatUint(res, 10))
}

///////////////////////////////////////////////////////////////////////////////
// Ackermann Section
func NaiveAckermannHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	m, err := strconv.ParseUint(vars["m"], 10, 64)
	if err != nil {
		if cfg.Debug == true {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "FIXME")
		}
	}
	n, err := strconv.ParseUint(vars["n"], 10, 64)
	if err != nil {
		if cfg.Debug == true {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "FIXME")
		}
	}
	// m and n have been acquired
	res, err := naiveAckermann(n, m)
	if err != nil {
		if cfg.Debug == true {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "FIXME")
		}

	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, strconv.FormatUint(res, 10))
	}
}

///////////////////////////////////////////////////////////////////////////////
// Factorial section
func NaiveFactorialHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	candidate, err := strconv.ParseUint(vars["n"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "FIXME")
	}
	res := naiveFactorial(candidate)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, strconv.FormatUint(res, 10))
}

func SmartFactorialHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	candidate, err := strconv.ParseUint(vars["n"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if cfg.Debug == true {
			fmt.Fprintf(w, err.Error())
		} else {
			fmt.Fprintf(w, "FIXME")
		}
	}
	res := smartFactorial(candidate)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, strconv.FormatUint(res, 10))
}

///////////////////////////////////////////////////////////////////////////////
// Home section
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

///////////////////////////////////////////////////////////////////////////////
// Help - if allowed
func HelpHandler(w http.ResponseWriter, r *http.Request) {
	msg := `
	endpoints: v2 is always more optimized than v1 and should be preferred
	/v1/fibonacci/{n:[0-9]+}
	/v2/fibonacci/{n:[0-9]+}
	/v1/ackermann/{n:[0|1|2|3]}/{m:[0-9]+}
	/v1/factorial/{n:[0-9]+}
	/v2/factorial/{n:[0-9]+}
	/status/memory
	/status/memory/raw
	/status/goroutines
	/debug/stackdump
	/debug/heapdump
	/debug/gcstats
	`
	if cfg.Debug == true {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, msg)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 page not found")
	}
}
