package main

import "testing"

func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func TestIntMin(t *testing.T) {
	ans:=IntMin(2,-2)
	if ans!=-2{
		t.Errorf("IntMin(2,-2)=%d;want -2",ans)
	}
}
