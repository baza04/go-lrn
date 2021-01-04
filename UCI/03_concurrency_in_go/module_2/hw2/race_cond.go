package main

import (
	"fmt"
	"time"
)

/*
In general, when each of the go routines is launched, it happens regardless of all the remaining executable processes.
by the way, the main function is launched in the same way, in our code we launched 10 go - routines, while the order of their end
may differ from the order of launch, which we observe. The reason it seems to me is that all independent go -
routines spend different time running the required packages in go and completing their tasks, and  Go Runtime Scheduler is impact.

Race conditions occur when the outcome of a program is not always the same each time it is run.
Race conditions arise because the end result depends on an undefined order in which the goroutines are executed.
As you can see in this program, the someFuns function requires a string and a number to be passed.
The problem starts when 10 goroutines are launched at once. Logically, the program should print
the line "first" with a number, then the line "second" and the number, but in reality, this is
not quite the case. But since the order of rotation is not deterministic, and we do not control
their execution either by a lock or a timer, it is likely that any of the 10 goroutines will be
executed before those that were launched much earlier.

You can see the race conditions when you run this program.
*/

func main() {
	for i := 0; i < 5; i++ {
		go someFunc("first", i)
		go someFunc("second", i)
	}
	time.Sleep((time.Millisecond * 100))
}

func someFunc(str string, i int) {
	fmt.Println(str, i)
}
