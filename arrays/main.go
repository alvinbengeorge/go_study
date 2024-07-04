package main

import "fmt"

func main() {
	var array [5]int
	fmt.Println(array, len(array), cap(array))
	array = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array, len(array), cap(array))

	var array2 = [2]string{"Hello", "World"}
	fmt.Println(array2, len(array2), cap(array2))

	array2[1] = "Golang"
	fmt.Println(array2, len(array2), cap(array2))

	var array3 []int
	fmt.Println(array3, len(array3), cap(array3))
	array3 = append(array3, 1)
	fmt.Println(array3, len(array3), cap(array3))
}
