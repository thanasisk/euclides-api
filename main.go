package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Debug        bool          `json:debug`
	ReadTimeout  time.Duration `json:readTimeout`
	WriteTimeout time.Duration `json:writeTimeout`
	IdleTimeout  time.Duration `json:idleTimeout`
	Port         string        `json:port`
	Address      string        `json:address`
}

var cfg Config

// init is located elsewhere too!
func init() {
	// first and foremost load configuration
	// FIXME - if there is time, make it 12-factor?
	cfg = loadConfiguration("config.json")
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

// an utility method to load JSON configuration
func loadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
