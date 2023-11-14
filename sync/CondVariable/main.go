package main

import (
	"fmt"
	"sync"
)

var sharedsrc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup

	mu := sync.Mutex{}
	c := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		c.L.Lock()
		for len(sharedsrc) < 1 {
			// time.Sleep(time.Second)
			c.Wait() // suspends the go routine
		}
		fmt.Println(sharedsrc["hii"])
		c.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		c.L.Lock()
		for len(sharedsrc) < 2 {
			// time.Sleep(time.Second)
			c.Wait()
		}
		fmt.Println(sharedsrc["hii__"])
		c.L.Unlock()

	}()

	c.L.Lock()
	sharedsrc["hii"] = "hii"
	sharedsrc["hii__"] = 45
	c.L.Unlock()
	c.Broadcast() // signals all stopped g routines
	wg.Wait()

}
