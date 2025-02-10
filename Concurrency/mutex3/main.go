package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var balance = 0

func incBy1() {
	defer wg.Done()
	balance += 1
}

func main() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go incBy1()
	}
	wg.Wait()
	fmt.Println(balance)
}
