package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
}

var wg sync.WaitGroup

func (p Philo) eat(number int, ch chan int) {
	for i := 0; i < 3; i++ {
		<-ch // get permission  from host
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("starting to eat %d\n", number)

		// it make possible to 2 people eat at the same time (not first start, first finish)
		time.Sleep(time.Millisecond * 100)

		ch <- number // return permission
		fmt.Printf("finish eating %d\n", number)
		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
	wg.Done()
}

func host(ch chan int) {
	wg.Add(1)
	for i := 0; i < 15; i++ {
		// give permission to eat (for 2 people at the same time))
		ch <- 1
		ch <- 2
		// take permission
		<-ch
		<-ch
	}
	wg.Done()
}

func initArrs(CSticks []*ChopS, philos []*Philo) {
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}
}

func main() {
	CSticks := make([]*ChopS, 5)
	philos := make([]*Philo, 5)
	initArrs(CSticks, philos)

	ch := make(chan int, 2)
	go host(ch)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(i+1, ch)
	}
	wg.Wait()
}
