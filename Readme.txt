so basically if you push data into a channel(A to B) from a go routine(A) that go routine
will be blocked until the same message is read from the go routine(B)

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
