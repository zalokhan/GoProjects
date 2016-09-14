package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func player(name string, c chan int) {
	// YOUR CODE!!
	defer wg.Done()
	for {
		unicorn, ok := <-c
		if !ok {
			return
		}
		fmt.Printf("<- [%s] received the Unicorn(%d)\n", name, unicorn)

		if rand.Intn(100)%21 == 0 {
			fmt.Printf("%s is B.O.R.E.D !!\n", name)
			close(c)
			return
		}
		unicorn++
		fmt.Printf("-> [%s] sending unicorn(%d)\n", name, unicorn)
		c <- unicorn
	}
}

func main() {
	u := make(chan int)

	wg.Add(2)
	go player("Norbet", u)
	go player("Faranouch", u)

	u <- 0

	wg.Wait()
}
