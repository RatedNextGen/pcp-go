package main

import "fmt"

func main() {

	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]
	fmt.Println("array: ", array) //Output -> array:  [1 2 3 4 5]
	fmt.Println("slice: ", slice) //Output -> slice:  [2 3 4]

	fmt.Println("len: ", len(slice)) //Output -> len:  3
	fmt.Println("cap: ", cap(slice)) //Output -> cap:  4

	slice = append(slice, 10)
	fmt.Println("newslice: ", slice) //Output -> newslice: [2 3 4 10]

	slice[0] = 100
	fmt.Println("array: ", array) //Output -> array:  [1 100 3 4 10]
	fmt.Println("slice: ", slice) //Output -> slice:  [100 3 4 10]

}
