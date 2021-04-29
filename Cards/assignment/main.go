package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range nums {
		if nums[index]%2 == 0 {
			fmt.Println(value, "even")
		} else {
			fmt.Println(value, "odd")
		}
	}
}
