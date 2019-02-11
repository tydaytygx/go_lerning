package main

import "fmt"

func main() {
	a := 10
	b := 20

	defer func(a, b int) {
		fmt.Println(a, b)
	}(a, b) //值在调用前已经传入，后续变量更改没有影响
	a = 100
	b = 200
	fmt.Println(a, b)
}
