package main

import (
	"fmt"
	"time"
)

const (
	ThreeSeconds = 3 * time.Second
	HalfSecond = 500 * time.Millisecond

)

func doBlockingWait(d time.Duration){
	time.Sleep(d)
}

func doDemo(){
	done := make(chan bool)

	go func(){
		doBlockingWait(ThreeSeconds)
		fmt.Print("1")
		result := "42"

		doBlockingWait(ThreeSeconds)
		fmt.Print("2")
		finalResult := "The answer is " + result

		doBlockingWait(ThreeSeconds)
		fmt.Print(finalResult)

		done <- true
	}()

	fmt.Println("-> Now waiting for things to happen!")

loop:
		for{
			select {
				case <- done:
					break loop
				default:
					fmt.Print(".")
					doBlockingWait(HalfSecond)
			}

		}
		fmt.Println("\n-> Done.")
}

func main(){
	doDemo()
}
