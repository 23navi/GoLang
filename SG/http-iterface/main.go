package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Body)

	io.Copy(os.Stdout, resp.Body)
}


func printSomething(value any){
	fmt.Println(value)
}