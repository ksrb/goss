//http://tour.golang.org/moretypes/22
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev := 0
	curr := 1
	total := 0
	return func() int {
		prev = curr
		curr = total
		total = prev + curr
		return total
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
