package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	i := 0
	j := 0

	for i < 5 {
		j = 0
		for j <= i {
			fmt.Print(i + 1)
			j++
		}
		fmt.Println()
		i++
	}

	var name = []string{"Alvin", "Ben", "George"}
	for i, name := range name {
		fmt.Println(i, name)
	}

}
