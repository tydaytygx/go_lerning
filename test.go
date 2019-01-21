package main

import "fmt"

func main() {
	a := make(map[int]string)
	a[1] = "suckoff"
	a[2] = "gaybar"
	for k, v := range a {
		fmt.Println("key=", k, "value=", v)
	}
}
