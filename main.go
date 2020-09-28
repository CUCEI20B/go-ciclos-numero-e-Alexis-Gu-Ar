package main

import "fmt"

func main() {
	var e float64 = 0
	var lim int

	fmt.Scan(&lim)

	for i := 0; i < lim; i++ {
		fact := 1.0
		for j := i; j > 0; j-- {
			fact *= float64(j)
		}
		e += (1.0 / fact)
	}
	if lim < 10 && lim > 2 {
		fmt.Printf("%.1f", e)
	} else {
		fmt.Print(e)
	}
}
