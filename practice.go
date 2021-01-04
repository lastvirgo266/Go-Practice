package main

import "fmt"

func fac(n int) int {
	if n<= 0 {
		return 1
	}
	return n* fac(n-1)
}

func main() {
	fmt.Println("Hello")

	var i = 10
	b := "little"
	b += "Gophers"
	fmt.Println("Hello", i, b)
	fmt.Println(fac(5))

}
