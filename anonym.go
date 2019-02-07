package main

import "fmt"

func main() {
	func(a string) {
		fmt.Println(a)
	}("你是gay")

	f := func(a string) {
		fmt.Println(a)
	}

	f("你是gay")

}
