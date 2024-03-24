package main

import "fmt"

type Store struct {
	Name      string
	Inventory map[string]int // maps from item names to quantity
}

type Transaction struct {
	ItemName  string
	Quantity  int // quantity can be positive or negative
}

func updateInventory(s *Store, t Transaction) {

}

func main() {
	s := &Store{
		Name: "Store 1",
		Inventory: map[string]int{
					 "Item 1": 10,
					 "Item 2": 20,
	}, }
	t := Transaction{
		ItemName: "Item 1",
		Quantity: -5,
	}
	updateInventory(s, t)
	fmt.Println("Updated store:")
	fmt.Println(*s)

}