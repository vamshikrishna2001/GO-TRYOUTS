package main

import "fmt"

func generator(num ...int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, i := range num {
			ch <- i
		}
		close(ch)

	}()
	return ch
}

func square(ch <-chan int) <-chan int {
	ch2 := make(chan int)
	go func() {
		for n := range ch {
			ch2 <- n * n
		}
		close(ch2)
	}()

	return ch2
}

func main() {
	fmt.Println("hii")
	for o := range square(square(generator(388, 26666))) {
		fmt.Println(o)
	}

}
