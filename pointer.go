package main

import "fmt"

func gaybar(a int) {
	a = 2
}

func gaybar2(a *int) {
	*a = 2
}

func main() {
	a := 23
	b := 23
	gaybar(a)
	gaybar2(&b)

	fmt.Println(a)
	fmt.Println(b)

}
