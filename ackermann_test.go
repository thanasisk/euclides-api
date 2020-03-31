package main

type ackermannTuple struct {
	m int
	n int
}

type ackermannTest struct {
	expected uint64
	m        int
	n        int
}

var ackermannTests []ackermannTest

func init() {
	ackermannTests = []ackermannTest{
		{1, 0, 1},
		{},
		{},
		{},
		{},
		{},
		{},
	}
}
