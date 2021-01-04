package main

import (
	"fmt"
	"time"
)

func Increment(value *int) {
	*value += 1
}

func PrintValue(value *int) {
	fmt.Println(*value)
}

func main() {
	for i := 0; i < 5; i++ {
		value := 1

		go Increment(&value)
		go PrintValue(&value)

		time.Sleep(time.Second)
	}
	fmt.Println("Done")
}

/*
Race Conditions occurs when the output of the program is not the same everytime you run it.
Race Conditions happens because the outcome depends on non-deterministic ordering.
As you can see in this program, the function Increment and PrintValue are running concurrently.
The problem starts when the two functions are using the same source (which in this case is the variable 'value') to perform different methods.
The real result for this program (without the methods running concurrently) is that the value of 'value' is increased by 1 and then printed in the command prompt for 5 times.
But because the order of interleavings is non-deterministic, there's a chance that the task in PrintValue is executed first before the task in Increment, or vice versa.
So in all 5 iterations, there's a possibility that one iteration prints the value as 1 instead of 2.
There's also a possibility that all 5 iterations print the value as 1 or as 2.
You can see the race conditions when you run this program.
*/
