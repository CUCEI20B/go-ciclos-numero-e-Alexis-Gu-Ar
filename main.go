package main

import "fmt"

func main() {
	var e float64
	var lim int

	fmt.Scan(&lim)

	for i := 0; i < lim; i++ {
		fact := 1
		for j := i; j > 0; j-- {
			fact *= j
		}
		e = e + (1.0 / float64(fact))
	}
	fmt.Print(e)
}
