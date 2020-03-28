package main

import "fmt"
import "github.com/gorilla/mux"
import "net/http"
import "time"
import "log"
import "strconv"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/v1/fibonacci/{n}", NaiveFibonacciHandler)
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
	candidate, err := strconv.Atoi(vars["n"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "FIXME")
	}
	res := naiveFibonacci(candidate)
	if res < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please provide only non-negative values")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, strconv.Itoa(res))
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

func naiveFibonacci(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return naiveFibonacci(n-2) + naiveFibonacci(n-1)
	}
}
