package main

import "fmt"

type Meter int
type Kilometer int

type Dog struct{}

func (Dog) makeNoise() {
	fmt.Println("wufwuf")
}

type Noise interface {
	makeNoise()
}

func main() {
	var noisemaker Noise = Dog{}
	noisemaker.makeNoise()

	var m Meter = 10
	var km Kilometer = 10
	fmt.Println(m, km)
}
