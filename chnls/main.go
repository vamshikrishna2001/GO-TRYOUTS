package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		t := <-ch
		fmt.Println(t)
	}()

	ch <- 2
	wg.Wait()
}
