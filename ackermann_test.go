package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type ackermannTuple struct {
	m int
	n int
}

type ackermannTest struct {
	n        int
	m        int
	expected int
}

var ackermannTests []ackermannTest

func init() {
	ackermannTests = []ackermannTest{
		{0, 1, 2},
		{0, 2, 3},
		{0, 3, 4},
		{1, 1, 3},
		{1, 2, 4},
		{1, 3, 5},
		{2, 1, 5},
		{2, 2, 7},
		{2, 3, 9},
		{3, 1, 13},
		{3, 2, 29},
		{3, 3, 61},
	}
}

func TestNaiveHandlerShouldPass(t *testing.T) {
	for _, tt := range ackermannTests {
		tPath := fmt.Sprintf("/v1/ackermann/%d/%d", tt.n, tt.m)
		req, err := http.NewRequest("GET", tPath, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc("/v1/ackermann/{n}/{m}", NaiveAckermannHandler)
		router.ServeHTTP(rr, req)

		expected := strconv.Itoa(tt.expected)
		if rr.Code == http.StatusOK && expected != rr.Body.String() {
			//
			t.Errorf("HTTP Code: %d got: %v want: %v",
				rr.Code, rr.Body.String(), expected)
		}
	}

}
