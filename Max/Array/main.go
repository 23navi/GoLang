package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%v %v %v %p \n", slice, len(slice), cap(slice), &slice[0])
	slice2 := append(slice, 11)
	slice = append(slice, slice2...)
	fmt.Printf("%v %v %v %p \n", slice, len(slice), cap(slice), &slice[0])

}
