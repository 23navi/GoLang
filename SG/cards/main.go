package main

import "io/ioutil"

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// const abc int = 20
	// fmt.Print(abc)

	ioutil.WriteFile("test.txt", []byte("Hello World"), 0666)

}
