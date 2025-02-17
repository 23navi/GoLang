package main

import (
	"fmt"

	"github.com/23navi/go-toolkit"
)

func main() {
	tools := toolkit.Tools{}
	randomString := tools.RandomString(10)
	fmt.Println(randomString)
}
