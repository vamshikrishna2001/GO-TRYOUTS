so basically if you push data in a channel from a go routine that go routine
will be blocked until the same message is read from the go routine

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