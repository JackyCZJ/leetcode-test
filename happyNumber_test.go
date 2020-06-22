package main

import "testing"

func Test_isHappy(t *testing.T) {
	if !isHappy(19) {
		t.Fatal("error")
	}
}