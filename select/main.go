package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(ch1 chan<- string) {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}(ch1)

	go func(ch2 chan<- string) {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}(ch2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-ch1:
			fmt.Println(m1)
		case m2 := <-ch2:
			fmt.Println(m2)
		}
	}

}
