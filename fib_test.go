package main

import "testing"

func TestFib(t *testing.T) {

	d:=fib()
	if d() != 1 {
		t.Errorf("value is not perfect %v",d())
	}

}