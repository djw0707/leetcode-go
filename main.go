package main

import (
	"fmt"
	"sync"
)

func main() {
	chan1 := make(chan int, 1)
	chan2 := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 5; i++ {
			<-chan2
			fmt.Println(i * 2)
			chan1 <- 1
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 5; i++ {
			<-chan1
			fmt.Println(i*2 + 1)
			chan2 <- 1
		}

	}()
	chan2 <- 1
	wg.Wait()
}
