package main

import "fmt"

func main() {

	var arr = []int{1, 2, 3, 4}
	var new = []int{}
	var sum int
	for i := 0; i < len(arr); i++ {
		sum = arr[i] + sum

		//fmt.Println(sum)
		new = append(new, sum)

	}
	fmt.Println(new)

}
