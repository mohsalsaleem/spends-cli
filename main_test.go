package main

import "testing"

func TestToFile(t *testing.T) {
	a, _ := toFils("200")

	if a != 20000 {
		t.Log(a)
		t.Fail()
	}
}
