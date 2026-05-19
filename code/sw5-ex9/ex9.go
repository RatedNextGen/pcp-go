package main

import "fmt"

func collatzIterative(n int) {
	for n > 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println()
}

func main() {
	fmt.Print("Folge für 3: ")
	collatzIterative(3)

	fmt.Print("Folge für 7: ")
	collatzIterative(7)
}
