package main

import (
	"fmt"
	"strings"
)

func main() {
	var hello = "Hello"
	var world = "World"
	var helloWorld = strings.Join([]string{hello, world}, " ")

	fmt.Println(strings.ToUpper(helloWorld))
}
