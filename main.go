package main

import "fmt"

func main() {
	var e float64 = 0
	var lim int

	fmt.Scan(&lim)

	for i := 0; i <= lim; i++ {
		var fact uint64 = 1
		for j := i; j > 0; j-- {
			fact *= uint64(j)
		}
		e += (1.0 / float64(fact))
	}

	fmt.Print(e)

}
