package main

import "fmt"

func gen(ch1 chan<- string) {

	ch1 <- "VA"
}

func relay(ch1 <-chan string, ch2 chan<- string) {
	temp := <-ch1

	ch2 <- temp
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go gen(ch1)

	go relay(ch1, ch2)

	fmt.Println(<-ch2)

}
