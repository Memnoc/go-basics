package main

import "fmt"

/*
*
* NOTE: first attempt
*
 */
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

/*
*
* NOTE: refactoring
*
 */
// func is_even_or_odd() {
// 	list := []int{7, 2, 9, 3, 6}
// 	for _, num := range list {
// 		if num%2 == 1 {
// 			fmt.Println("Number is odd")
// 		} else {
// 			fmt.Println("Number is even")
// 		}
// 	}
// }

/*
*
* NOTE: refactoring three - more flexibility for testing
*
 */

func is_even_or_odd_two(nums []int) string {
	result := ""
	for _, num := range nums {
		if num%2 == 1 {
			result = "odd"
			fmt.Println("Number is", result)
		} else {
			result = "even"
			fmt.Println("Number is", result)
		}
	}
	return result
}
