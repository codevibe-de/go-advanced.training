package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		fmt.Println("Hello from your Goroutine!")
		wg.Done()
	}()
	wg.Wait()
}
