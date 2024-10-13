package main

import (
	"first/doctor"
	"fmt"
)

func main() {
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)
}
