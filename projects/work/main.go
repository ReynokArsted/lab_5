package main

import (
	"fmt"
	"sync"
	"time"
)

func work() {
	time.Sleep(time.Millisecond * 50)
	fmt.Println("done")
}

func main() {
	wg := new(sync.WaitGroup)
	var counter int = 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
			counter += 1
		}()
	}
	wg.Wait()
	fmt.Printf("Кол-во горутин, которые закончили work() равно: %d\n", counter)
}
