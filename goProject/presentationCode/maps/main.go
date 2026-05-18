package main

import "fmt"

type Order struct {
	ID int
	Product string
	Quantity int
	Status string
}

func printOrders(orders []Order){
	for _, order := range orders {
		fmt.Printf("Bestellung %d: %s x%d [%s]\n", 
			order.ID, order.Product, order.Quantity, order.Status)
	}
}

func buildOrderIndex(orders []Order) map[int]Order {
	index := make(map[int]Order)

	for _, order := range orders {
		index[order.ID] = order
	}
	return index
}

func main(){
	orders := []Order{
		{ID: 1, Product: "Kaffee", Quantity: 2, Status: "neu"},
		{ID: 2, Product: "Tee", Quantity: 1, Status: "neu"},
		{ID: 3, Product: "Wasser", Quantity: 4, Status: "neu"},
	}

	printOrders(orders)

	orderIndex := buildOrderIndex(orders)
	fmt.Println("\nDirekter Zugriff über Map:")
	order, exists := orderIndex[2]
	if exists {
		fmt.Printf("Gefunden: Bestellung %d = %s x%d [%s]\n",
			order.ID, order.Product, order.Quantity, order.Status)
	}
}
