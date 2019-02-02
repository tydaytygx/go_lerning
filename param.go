package main

import "fmt"

func add(args ...int) (result int) {

	for _, arg := range args {
		result += arg
	}
	return
}

func main() {
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
