package main

import (
	"fmt"
	"sync"
	"time"
)

var finishDine chan *philosopher = make(chan *philosopher)  // for inform finish dine
var requestDine chan *philosopher = make(chan *philosopher) // for request dine

var philoNum int = 5
var mealsNum int = 3

var eatWgroup sync.WaitGroup

type cs struct{ sync.Mutex }

type philosopher struct {
	id              int
	leftCs, rightCs *cs
	allowedIn       chan bool
	mealsHad        int
}
type Host struct {
	activeDiners int
	mealsServed  int
}

func (h *Host) serve() {
	fmt.Println("Host ready to serve...")
	for {
		select {

		case p := <-finishDine:
			fmt.Println("Host continues serving as philosopher", p.id, "finished dine number #", p.mealsHad)
			h.activeDiners--
			h.mealsServed++ //meals served

		case p := <-requestDine:
			if h.activeDiners < 2 {
				fmt.Println("Host helped philosopher", p.id, "to sit down.")
				h.activeDiners++
				p.allowedIn <- true
			} else {
				fmt.Println("Host is waiting for someone philosopher")
				time.Sleep(1 * time.Second)
				go func() { requestDine <- p }() // hack: push the request back
			}
		}

		// Enough meals were served to stop the dinner
		if h.mealsServed == (philoNum * mealsNum) {
			fmt.Println("The host finishes his work")
			eatWgroup.Done()
			return
		}
	}
}

func (p *philosopher) eat() {

	if p.mealsHad < mealsNum {
		//request dine
		requestDine <- p
		<-p.allowedIn
		p.leftCs.Lock()
		p.rightCs.Lock()

		say("starting to eat", p.id)
		time.Sleep(1 * time.Second)
		//increasing meals had
		p.mealsHad++
		p.rightCs.Unlock()
		p.leftCs.Unlock()

		say("finishing eating", p.id)
		finishDine <- p
		p.eat()
	} else {
		defer eatWgroup.Done()
		fmt.Println("Philosopher", p.id, "had finished their meals for the day")
	}

}

func say(action string, id int) {
	fmt.Printf("Philosopher #%d is %s\n", id+1, action)
}

func main() {
	//philosophers and forks(cs)
	count := 5

	// Create forks
	forks := make([]*cs, count)
	for i := 0; i < count; i++ {
		forks[i] = new(cs)
	}

	// Create philospoher, assign them 2 forks(chopsticks) and send them to the dining table
	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{id: i, leftCs: forks[i], rightCs: forks[(i+1)%count], allowedIn: make(chan bool)}
		eatWgroup.Add(1)

		go philosophers[i].eat()
	}
	// Init: the dinner host
	host := new(Host)
	eatWgroup.Add(1)
	go host.serve()

	eatWgroup.Wait()

}
