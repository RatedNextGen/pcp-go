package main

import "fmt"

type OrderID int

type Order struct {
	ID OrderID
	Product string
	Quantity int
	Status string
}

type Process interface {
	Process(order *Order)
}

type FastProcess struct{}
type SafeProcess struct{}

func (FastProcess) Process(order *Order){
	order.Status = "schnell verarbeitet"
}

func (SafeProcess) Process(order *Order){
	order.Status = "sicher verarbeitet"
}

func runProcess(p Process, order *Order){
	p.Process(order)
}

func main(){
	order := Order{
		ID: 1,
		Product: "Kaffee",
		Quantity: 2,
		Status: "neu",
	}

	fast := FastProcess{}
	runProcess(fast, &order)

	fmt.Printf("Bestellung %d: %s [%s]\n", order.ID, order.Product, order.Status)
}
