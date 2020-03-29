package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type fibTest struct {
	n        int
	expected int
}

var fibTests []fibTest

func init() {
	fibTests = []fibTest{
		{1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13}, {8, 21},
		{9, 34}, {10, 55}, {11, 89}, {12, 144}, {13, 233}, {14, 377}, {15, 610},
		{16, 987}, {17, 1597}, {18, 2584}, {19, 4181}, {20, 6765},
	}
}

// FIXME: code duplication
func TestNaiveFibonacciHandlerShouldPass(t *testing.T) {
	for _, tt := range fibTests {
		tPath := fmt.Sprintf("/v1/fibonacci/%d", tt.n)
		req, err := http.NewRequest("GET", tPath, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc("/v1/fibonacci/{n}", NaiveFibonacciHandler)
		router.ServeHTTP(rr, req)

		expected := strconv.Itoa(tt.expected)
		if rr.Code == http.StatusOK && expected != rr.Body.String() {
			//
			t.Errorf("HTTP Code: %d got: %v want: %v",
				rr.Code, rr.Body.String(), expected)
		}
	}
}

// FIXME: code duplication
func TestSmartFibonacciHandlerShouldPass(t *testing.T) {
	for _, tt := range fibTests {
		tPath := fmt.Sprintf("/v2/fibonacci/%d", tt.n)
		req, err := http.NewRequest("GET", tPath, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc("/v2/fibonacci/{n}", SmartFibonacciHandler)
		router.ServeHTTP(rr, req)

		expected := strconv.Itoa(tt.expected)
		if rr.Code != http.StatusOK {
			t.Errorf("HTTP Code: got: %d wanted: %d", rr.Code, http.StatusOK)
		}
		if rr.Code == http.StatusOK && expected != rr.Body.String() {
			//
			t.Errorf("HTTP Code: %d got: %v want: %v",
				rr.Code, rr.Body.String(), expected)
		}
	}
}
