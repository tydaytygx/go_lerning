package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Printf("%T\n", add)
}
