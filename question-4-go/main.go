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

func updateInventory(s *Store, t Transaction) bool {
	itemQuantity, itemExists := s.Inventory[t.ItemName]

	// Add Item
	if t.Quantity > 0 {
		if !itemExists {
			s.Inventory[t.ItemName] = t.Quantity
		} else {
			s.Inventory[t.ItemName] += t.Quantity
		}
		return true
	}

	// Delete Item
	if t.Quantity < 0 {
		if !itemExists || itemQuantity + t.Quantity < 0 {
			return false
		}
		s.Inventory[t.ItemName] += t.Quantity
		return true
	}

	return false
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
	result := updateInventory(s, t)
	fmt.Println(result)
	if result {
		fmt.Println("Store after execution:")
		fmt.Println(*s)
	} 

	s= &Store{
		Name: "Local Store",
		Inventory: map[string]int{
					 "Apples":  100,
					 "Bananas": 50,
					 "Oranges": 30,
		},
	}
	t = Transaction{
			ItemName: "Bananas",
			Quantity: -60,
	}
	result = updateInventory(s, t)
	fmt.Println(result)
	if result {
		fmt.Println("Store after execution:")
		fmt.Println(*s)
	} 

	s = &Store{
		Name: "Store 2",
		Inventory: map[string]int{
					 "Item 1": 15,
					 "Item 2": 25,
	}, }
	t = Transaction{
			ItemName: "Item 3",
			Quantity: 10,
	}
	result = updateInventory(s, t)
	fmt.Println(result)
	if result {
		fmt.Println("Store after execution:")
		fmt.Println(*s)
	} 

	s = &Store{
		Name: "Store 3",
		Inventory: map[string]int{
					"Item 1": 10,
					"Item 2": 5,
		},
	}
	t = Transaction{
		ItemName: "Item 2",
		Quantity: -10,
	}
	result = updateInventory(s, t)
	fmt.Println(result)
	if result {
		fmt.Println("Store after execution:")
		fmt.Println(*s)
	} 

	s = &Store{
		Name: "Store 4",
		Inventory: map[string]int{
					 "Item 1": 5,
	}, }
	t = Transaction{
			ItemName: "Item 1",
			Quantity: -10,
	}
	result = updateInventory(s, t)
	fmt.Println(result)
	if result {
		fmt.Println("Store after execution:")
		fmt.Println(*s)
	} 

	s = &Store{
		Name: "Store 5",
		Inventory: map[string]int{
					 "Item 1": 10,
					 "Item 2": 5,
		},
	}
	t = Transaction{
			ItemName: "Item 2",
			Quantity: 20,
	}
	result = updateInventory(s, t)
	fmt.Println(result)
	if result {
		fmt.Println("Store after execution:")
		fmt.Println(*s)
	} 

	s = &Store{
		Name: "Store 6",
		Inventory: map[string]int{},
	}
	t = Transaction{
				ItemName: "Item 1",
				Quantity: 5,
	}
	result = updateInventory(s, t)
	fmt.Println(result)
	if result {
		fmt.Println("Store after execution:")
		fmt.Println(*s)
	} 
}