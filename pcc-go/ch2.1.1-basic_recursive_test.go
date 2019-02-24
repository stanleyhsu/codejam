package main

import "testing"

var ValidTable = map[int]int{
	1: 1,
	2: 2,
	3: 6,
	5: 120,
}

func TestFact0(t *testing.T) {
	if 1 != fact(0) {
		t.Errorf("fact 0 not success")
	}
}

func TestFactNormal(t *testing.T) {
	for k, v := range ValidTable {
		if v != fact(k) {
			t.Errorf("fact(%v) Expected(%v) but Got(%v)", k, v, fact(k))
		}
	}
}
