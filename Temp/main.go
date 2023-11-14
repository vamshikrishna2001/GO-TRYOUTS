package main

import (
	"fmt"
	"time"
)

// func main_1() {
// 	// var out int

// 	ch := make(chan int)

// 	go func() {
// 		time.Sleep(10 * time.Second)
// 		i := 1000
// 		ch <- i

// 	}()

// 	temp := <-ch
// 	fmt.Println(temp)

// 	fmt.Println("hii")

// }

// func gen() chan int {
// 	ch := make(chan int)

// 	go func() {
// 		for i := 0; i < 4; i++ {
// 			ch <- i
// 		}
// 		close(ch)

// 	}()
// 	return ch

// }

// func main() {

// 	ch := gen()

// 	for i := range ch {
// 		fmt.Println(i)
// 	}

// }

// func main() {
// 	var a string
// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go func(a *string) {
// 		defer wg.Done()
// 		time.Sleep(6 * time.Second)
// 		*a = "vamshi"
// 	}(&a)

// 	wg.Wait()
// 	fmt.Println(a)

// }

func main() {
	var a int
	ch := make(chan int)

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("in go routine")
		a = <-ch

	}()
	ch <- 10

	fmt.Println(a)

}
