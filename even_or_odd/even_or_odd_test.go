package main

import (
	"fmt"
	"testing"
)

func TestEvenOrOdd(t *testing.T) {
	test := []int{6}
	result := is_even_or_odd_two(test)
	fmt.Println(result)

	if result != "even" {
		t.Errorf("Got odd expected, even %v", result)
	}

	fmt.Println(result)
}
