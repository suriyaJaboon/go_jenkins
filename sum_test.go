package main

import "testing"

func TestSum(t *testing.T) {
	s := Sum(1)
	if s != 2 {
		t.Errorf("Failure %d", s)
	}
}