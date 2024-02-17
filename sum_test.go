package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5.")
	}
}

func TestSub(t *testing.T) {

	result := sub(7, 4)

	if result != 3 {
		t.Error("The result must be 3.")

	}
}

func TestTimes(t *testing.T) {

	result := times(3, 4)

	if result != 12 {
		t.Error("The result must be 12.")
	}
}

func TestSumX(t *testing.T) {

	result := sumX(2, 5)

	if result != 9 {
		t.Error("The result must be 9.")
	}
}
