package main

import "testing"

func TestSum(t *testing.T) {
	x := sum (2,7)

	if x != 5{
		t.Error("Excepted", 5, "Got", x)
	}

}
