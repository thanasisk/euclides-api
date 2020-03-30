package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type factTest struct {
	n        int
	expected int
}

var factTests []factTest

func init() {
	factTests = []factTest{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
		{8, 40320},
		{9, 362880},
		{10, 3628800},
	}
}

func TestNaiveFactorialHandlerShouldPass(t *testing.T) {
	for _, tt := range factTests {
		tPath := fmt.Sprintf("/v1/factorial/%d", tt.n)
		req, err := http.NewRequest("GET", tPath, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc("/v1/factorial/{n}", NaiveFactorialHandler)
		router.ServeHTTP(rr, req)

		expected := strconv.Itoa(tt.expected)
		if rr.Code == http.StatusOK && expected != rr.Body.String() {
			//
			t.Errorf("HTTP Code: %d got: %v want: %v",
				rr.Code, rr.Body.String(), expected)
		}
	}

}

func TestSmartFactorialHandlerShouldPass(t *testing.T) {
	for _, tt := range factTests {
		tPath := fmt.Sprintf("/v2/factorial/%d", tt.n)
		req, err := http.NewRequest("GET", tPath, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc("/v2/factorial/{n}", SmartFactorialHandler)
		router.ServeHTTP(rr, req)

		expected := strconv.Itoa(tt.expected)
		if rr.Code == http.StatusOK && expected != rr.Body.String() {
			//
			t.Errorf("HTTP Code: %d got: %v want: %v",
				rr.Code, rr.Body.String(), expected)
		}
	}

}
