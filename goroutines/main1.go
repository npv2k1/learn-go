package main

import (
	"fmt"
	"sync"
)

func threadOne(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Printf("Thread 1: %d\n", i)
	}
}

func threadTwo(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Printf("Thread 2: %d\n", i)
	}

}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start threads")
	go threadOne(&wg)
	go threadTwo(&wg)
	fmt.Println("Waiting for threads to finish...")
	wg.Wait()
	fmt.Println("All threads finished")
}
