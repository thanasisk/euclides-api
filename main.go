package main

import "fmt"
import "github.com/gorilla/mux"
import "net/http"
import "time"
import "log"
import "strconv"

// init is located elsewhere ...

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/v1/fibonacci/{n}", NaiveFibonacciHandler)
	r.HandleFunc("/v2/fibonacci/{n}", SmartFibonacciHandler)
	r.HandleFunc("/v1/ackermann", NaiveAckermannHandler)
	r.HandleFunc("/v1/factorial", NaiveFactorialHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
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
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
