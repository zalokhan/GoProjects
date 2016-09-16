/**
The Unicorn Game. Create a program where players get to exchange unicorns over an
unbuffered channel. The game ends once either of the parties gets "randomly" seek and
tired of playing that game ;-).
**/
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
	defer wg.Done()

	for {
		unicorn, ok := <-c
		if !ok {
			return
		}
		fmt.Printf("<- [%s] received Unicorn(%d)\n", name, unicorn)

		if rand.Intn(100)%21 == 0 {
			fmt.Printf("%s is B.O.R.E.D!!\n", name)
			close(c)
			return
		}

		unicorn++
		fmt.Printf("-> [%s] sending Unicorn(%d)\n", name, unicorn)
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
