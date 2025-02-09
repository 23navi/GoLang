package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		if key != 0 {
			fmt.Println("You pressed", char, key)
		} else {
			fmt.Println("You pressed", char, key)
		}

		if key == keyboard.KeyEsc {
			break
		}
	}
}
