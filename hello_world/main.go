package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var x int16 = 5
	var y int16 = 7

	var sum int16 = x + y

	fmt.Println(x, "+", y, "=", sum)

	var a float32 = 5.5
	var b float32 = 7.7

	var sum2 float32 = a + b

	fmt.Println(a, "+", b, "=", sum2)

	var first_name string = "Alvin"
	var last_name string = "George"

	var full_name string = first_name + " " + last_name

	fmt.Println("My name is", full_name)

}
