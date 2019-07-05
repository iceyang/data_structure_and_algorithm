package main

import "testing"

func TestKnapsack(t *testing.T) {
	if knapsack(8, []int{2, 3, 4, 5}, []int{3, 4, 5, 6}) != 10 {
		t.Fail()
	}
}
