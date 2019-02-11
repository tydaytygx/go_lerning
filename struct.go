package main

import "fmt"

func main() {
	type gaybar struct {
		Name string
		Age  int
	}

	van := gaybar{
		Name: "Van DarkHolme",
		Age:  46, //注意需要逗号
	}
	fmt.Println(van)
}
