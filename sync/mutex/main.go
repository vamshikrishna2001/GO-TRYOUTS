package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	var counter int64
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				mutex.Lock()
				counter++ // or you can use atomic.add(counter ,1)
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
