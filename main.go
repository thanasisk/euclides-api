package main

import "fmt"
import "github.com/gorilla/mux"
import "net/http"
import "time"
import "log"
import "strconv"

// init is located elsewhere ...
func main() {
	logo := `
  ______           _ _     _
 |  ____|         | (_)   | |
 | |__  _   _  ___| |_  __| | ___  ___
 |  __|| | | |/ __| | |/ _  |/ _ \/ __|
 | |___| |_| | (__| | | (_| |  __/\__ \
 |______\__,_|\___|_|_|\__,_|\___||___/
`
	fmt.Print(logo)
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/v1/fibonacci/{n:[0-9]+}", NaiveFibonacciHandler)
	r.HandleFunc("/v2/fibonacci/{n:[0-9]+}", SmartFibonacciHandler)
	r.HandleFunc("/v1/ackermann", NaiveAckermannHandler)
	r.HandleFunc("/v1/factorial/{n:[0-9]+}", NaiveFactorialHandler)
	r.HandleFunc("/v2/factorial/{n:[0-9]+}", SmartFactorialHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

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

// to be revisited
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

func NaiveAckermannHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

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
		fmt.Fprintf(w, "FIXME")
	}
	res := smartFactorial(candidate)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, strconv.FormatUint(res, 10))
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
