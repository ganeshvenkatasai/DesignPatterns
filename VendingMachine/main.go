package main

import (
	"VendingMachine/inventory"
	"VendingMachine/machine"
	"VendingMachine/model"
)

func main() {
	inventory := inventory.NewInventory()
	vendingMachine := machine.NewVendingMachine(inventory)
	product := model.Product{Name: "Lays", Price: 20}

	vendingMachine.AddProduct(product)  // add product
	vendingMachine.PressButton(product) // accessing product without money
	vendingMachine.InsertCoin(10)       // inserting coin
	vendingMachine.InsertCoin(30)       // again inserting coin
	vendingMachine.PressButton(product) // accessing product
}
