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

func chocolateWorker(name string, coco <-chan int) {
	// YOUR CODE HERE!
	defer wg.Done()
	for {
		cocoNumber, ok := <-coco
		if !ok {
			return
		}
		fmt.Printf("[%s] wraps chocolate_%d\n", name, cocoNumber)
		time.Sleep(time.Duration(cocoNumber) * time.Second)
	}
}

func producer(coco chan<- int) {
	// YOUR CODE HERE!
	var cocoNumber int
	for {
		coco <- cocoNumber
		cocoNumber++
	}
}

func main() {
	coco := make(chan int, 10)

	wg.Add(2)
	go chocolateWorker("Lucy", coco)
	go chocolateWorker("Ethel", coco)

	producer(coco)

	wg.Wait()
}
