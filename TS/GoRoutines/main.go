package main

import (
	"fmt"
	"time"
)

func printSomething(something string) {
	fmt.Println(something)
}

func main() {

	printSomething("One")
	go printSomething("Two")
	time.Sleep(1 * time.Second)
	printSomething("Three")
}
