package main

import "testing"

func TestExample(t *testing.T) {
	result := returnFalse()
	if result {
		t.Fatal("failed test")
	}
}
