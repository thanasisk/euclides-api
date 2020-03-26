package main

import "fmt"
import "github.com/gorilla/mux"
import "net/http"
import "time"
import "log"

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
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, naiveFibonacci(vars["n"]))
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

func naiveFibonacci(n string) {

}
