package main

import "fmt"

func is_even_or_odd() {
	list := []int{7, 2, 9, 3, 6}
	for i := 0; i < len(list); i++ {
		if list[i]%2 == 1 {
			fmt.Println("Number is odd")
		} else {
			fmt.Println("Number is even")
		}
	}
}
