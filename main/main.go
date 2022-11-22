package main

import (
	"star"
)

func main() {
	// test regexp

	/* 	var n int
	   	fmt.Scanln(&n)
	   	students := make([]string, 10)
	   	for i := 0; i < n; i++ {
	   		fmt.Scanln(&students[i])
	   	}
	   	for i := 0; i < n; i++ {
	   		fmt.Println(students[i])
	   	}
	   	fmt.Println("-----answer-----")
	   	for i := 0; i < n; i++ {
	   		if star.IsValid(students[i]) {
	   			star.GetAnswer(students[i])
	   		} else {
	   			fmt.Println("False")
	   		}
	   	} */

	// test bubbleSort
	/* 	test := []int{2, 3, 1, 4, 2, 0, 1, 10, -1, -3, -10, 0, 239}
	   	star.BubbleSort(test)
	   	fmt.Println(test) */
	star.Crud()
}
