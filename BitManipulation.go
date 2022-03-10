package main

import (
	"fmt"
)

func main() {
	n1 := add(3, 10)
	fmt.Printf("%0b = %d\n", n1, n1)
	n2 := subtract(13, 4)
	fmt.Printf("%0b = %d\n", n2, n2)
} 

func add(x, y int) int {
	keep := (x & y) << 1
    res := x^y
 
    if (keep == 0) {
        return res
	}

    return add(keep, res)
}

func subtract(x, y int) int {
	return add(x, (^y +1))
}