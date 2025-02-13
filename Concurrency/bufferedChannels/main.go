package main

import (
	"fmt"
)

func worker(ch chan string) {
	for {
		data := <-ch

		for i := 0; i < 1000000; i++ {
			for i := 0; i < 1000000; i++ {
			
			}
		}
		fmt.Printf("%s is processed\n", data)
	}
}

func main() {
	buffChan := make(chan string, 10)
	for i := 0; i < 1000; i++ {
		go worker(buffChan)
	}

	for i := 0; i < 100000; i++ {
		buffChan <- fmt.Sprintf("%d", i)
		fmt.Printf("Message %d sent to the channel\n", i)
	}
	fmt.Println("Done processing of 100 messages")
	close(buffChan)
}
