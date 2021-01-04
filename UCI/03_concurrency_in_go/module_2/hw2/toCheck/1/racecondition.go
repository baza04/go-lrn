package main

import (
	"fmt"
	"sync"
)

const nb = 10000000

func routine1(i *int64, wg *sync.WaitGroup) {
	for j := 0; j < nb; j++ {
		(*i)++
	}
	wg.Done()
}
func routine2(i *int64, wg *sync.WaitGroup) {
	for j := 0; j < nb; j++ {
		(*i)++
	}
	wg.Done()
}
func main() {
	var i int64 = 0
	var wg sync.WaitGroup
	wg.Add(2)
	go routine1(&i, &wg)
	go routine2(&i, &wg)

	wg.Wait()
	fmt.Printf("%d != %d", i, nb*2)
}

/* in this example both goroutine access the same variable
at the same time, because of this they modify the variable
at the same time, and the number of incrementation is not
equal to the number (and not consistant) 


в этом примере обе горутины обращаются к одной и той же 
переменной в одно и то же время, из-за этого они изменяют 
переменную в одно и то же время, и количество приращений 
не равно числу (и не согласовано)
*/
