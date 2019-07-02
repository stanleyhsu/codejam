package main

import (
	"testing"
)

func TestFact0(t *testing.T) {
	if 1 != fact(0) {
		t.Error("fact 0 not success")
	}
}

var ValidTable = map[int]int{
	1: 1,
	2: 2,
	3: 6,
	5: 120,
}

func TestFactNormal(t *testing.T) {
	for k, v := range ValidTable {
		if v != fact(k) {
			t.Errorf("fact(%v) Expected(%v) but Got(%v)", k, v, fact(k))
		}
	}
}

func TestFib0(t *testing.T) {
	if 0 != fib(0) {
		t.Error("fib 0 not success")
	}
}

func TestFib1(t *testing.T) {
	if 1 != fib(1) {
		t.Error("fib 0 not success")
	}
}

var FibValidTable = map[int]int{
	0: 0,
	1: 1,
	2: 1,
	3: 2,
	4: 3,
	5: 5,
	6: 8,
	7: 13,
}

func TestFibNormal(t *testing.T) {
	for k, v := range FibValidTable {
		if v != fib(k) {
			t.Errorf("fib(%v) = %v, Expeted %v", k, fib(k), v)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	fib(30)
}

func BenchmarkFibm(b *testing.B) {
	fibm(30)
}
