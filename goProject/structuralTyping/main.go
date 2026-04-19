package main

import (
	"fmt"
	"structuralTyping/animals"
	"structuralTyping/robots"
	"structuralTyping/noises"
)

func Announce(n noises.Noisemaker) {
	fmt.Println(n.MakeNoise())
}

func main(){
	dog := animals.Dog{Name: "Dog"}
	robot := robots.Robot{Model: "X35"}

	Announce(dog)
	Announce(robot)
}
