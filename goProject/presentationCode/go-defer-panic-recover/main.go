package main

import "fmt"

type Order struct {
	ID       int
	Product  string
	Quantity int
	Status   string
}

func buildOrderIndex(orders []Order) map[int]Order {
	if len(orders) == 0 {
		panic("Kritischer Fehler: Bestellliste ist leer!")
	}
	index := make(map[int]Order)
	for _, order := range orders {
		index[order.ID] = order
	}
	return index
}

func main() {
	defer fmt.Println("Ich werde zuletzt ausgeführt!")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[Recover] Fataler Fehler abgefangen:", r)
		}
	}()

	var emptyOrders []Order
	buildOrderIndex(emptyOrders)
}
