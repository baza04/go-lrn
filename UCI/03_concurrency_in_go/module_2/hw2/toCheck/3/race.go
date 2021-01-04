/**
Race conditions are where 2 threads are accessing memory at the same time,
one of which is writing. Race conditions occur because of unsynchronized access to shared memory.
An interleaving of execution steps creates an incorrect result.
The issue is that multiple go routines are overwriting the same value.

In the following program, we are using a common shared variable which is being accesible by the two goroutines.
As the execution of goroutines will be in non-deterministic order, we can't predict the output value.
Output might be different at each time we execute the program.
*/

package main

import (
	"fmt"
	"time"
)

var number int

func race() {
	fmt.Println(number)
	number++
}

func main() {

	number = 1
	go race()
	go race()
	fmt.Println("Race condition when trying to access the same function")
	time.Sleep(1 * time.Second)
}
