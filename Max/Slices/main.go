package main

import "fmt"

func main() {
	myArr := [5]int{}
	fmt.Printf("%p\n", &myArr)
	slice := myArr[:]
	fmt.Printf("%p\n", slice)
	myFunc(&slice)

}

func myFunc(slice *[]int) {
	fmt.Printf("%p\n", slice)
}
