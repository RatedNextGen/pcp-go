package main

import "fmt"

func goDefer() {
	defer fmt.Println("Executed at the end")
	fmt.Println("Start ...")
	//Output:
	//Start ...
	//Executed at the end
}

func goPanic() {
	fmt.Println("Befor panic ")
	panic("error XY")
	fmt.Println("Will not be executed")
	//Output:
	//Befor panic
	//panic: error XY
}

func goRecover() {
	defer func() {
		error := recover()

		if error != nil {
			fmt.Println("Catch Error")
		}
	}()
	panic("Error XY")
	//Output:
	//Catch Error
}

func main() {

	goDefer()

	// goPanic()

	goRecover()

}
