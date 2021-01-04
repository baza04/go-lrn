package main

import (
	"fmt"
	"sync"
)

/* var i int

// var wg sync.WaitGroup
// var mx sync.Mutex
var ini sync.Once

func main() {

	// c := make(chan int, 2)
	// wg.Add(2)
	go someFunc()
	go someFunc()
	// for i := 0; i < 4; i++ {
	// 	x := <- c
	// 	fmt.Printf("c out: %d\n", x)
	// }
	// wg.Wait()
	time.Sleep(time.Millisecond * 100)
	fmt.Println(i)
}

func someFunc() {
	// for i := 0; i < 4; i++ {
	// 	c <- i
	// }
	// fmt.Printf("routine #%d\n", 4)
	// mx.Lock()
	i++
	// wg.Done()
	// mx.Unlock()
} */

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go dostuff()
	go dostuff()
	wg.Wait()
}

var on sync.Once

func setup() {
	fmt.Println("Init")
}

func dostuff() {
	on.Do(setup)
	fmt.Println("Stuff")
	wg.Done()
}
