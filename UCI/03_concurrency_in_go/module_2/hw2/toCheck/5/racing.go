package main

import (
	"fmt"
)

func calc(x *int) {
	*x++
}

func print(x *int) {
	fmt.Println(*x)
}

/*
racing conditions
The new functions calc and print are running in concurrent goroutines (threats).
Both functions are utilizing the "global" variable x.
Calc is incrementing x where print is printing.

The outcome of the this program is non deterministic, because it depends on the order of execution.
It is not possible to foresee the outcome.
Running this program multiple times will lead to no output at all, a few numbers or a lot of numbers in random order
*/
func main() {
	x := 0
	for i := 0; i < 100; i++ {
		go calc(&x)
		go print(&x)
	}
}
