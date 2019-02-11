package main

import "fmt"

func main() {
	func(a string) {
		fmt.Println(a)
	}("你是gay") //第一种方式

	f := func(a string) {
		fmt.Println(a)
	}

	f("你是gay") //第二种方式

}
