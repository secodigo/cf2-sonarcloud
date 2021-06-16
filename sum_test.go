package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(3, 2)

	if result != 5 {
		t.Error("The result must be five")
	}
}
