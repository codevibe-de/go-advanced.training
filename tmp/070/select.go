package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	chOfInts := make(chan int)
	chOfStrings := make(chan string)
	go worker(chOfStrings, chOfInts)
	chOfStrings <- "123"
	chOfStrings <- "222"
	_ = <-chOfInts
	_ = <-chOfInts
	time.Sleep(time.Second)
}

func worker(inCh chan string, outCh chan int) {
	var n int
	for {
		select {
		case s, ok := <-inCh:
			if ok {
				fmt.Println("Received ", s)
				n, _ = strconv.Atoi(s)
			}
		case outCh <- n:
			fmt.Println("Wrote ", n)
		}
	}
}
