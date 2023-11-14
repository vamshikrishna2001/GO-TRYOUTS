package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// func main() {
// 	wg.Add(1)
// 	go func(a, b int) {
// 		defer wg.Done()
// 		c := a + b
// 		fmt.Println(c, "from go routine")
// 	}(1, 2)

// 	wg.Wait()
// }

func main() {
	ch := make(chan int)
	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)

	r := <-ch
	fmt.Println(r)
}
