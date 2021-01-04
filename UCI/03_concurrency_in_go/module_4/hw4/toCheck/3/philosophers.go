package main

import (
	"fmt"
	"sync"
)

const (
	philCount = 5
	eatCount = 3
)

type cs struct {
	sync.Mutex
}

type phil struct {
	id              int
	leftCs, rightCs *cs
}

var wg sync.WaitGroup
var css = make([]*cs, philCount)
var phils = make([]*phil, philCount)

func init() {
	for i := 0; i < philCount; i++ {
		css[i] = new(cs)
	}

	for i := 0; i < philCount; i++ {
		phils[i] = &phil{
			id: i+1,
			leftCs: css[i],
			rightCs: css[(i+1)%philCount],
		}
	}
}

func main() {
	for _, phil := range phils {
		wg.Add(1)
		go phil.eat()
	}

	wg.Wait()
}

func (p phil) eat() {
	for j := 0; j < eatCount; j++ {
		p.leftCs.Lock()
		p.rightCs.Lock()

		fmt.Printf("starting to eat %d\n", p.id)
		fmt.Printf("finishing eating %d\n", p.id)

		p.leftCs.Unlock()
		p.rightCs.Unlock()
	}

	wg.Done()
}