package state

import (
	"Vendingmachine/inventory"
	"Vendingmachine/model"
)

type State interface {
	InsertCoin() error
	PressButton(amount int, p model.Product, i inventory.Inventory) error
	Dispense(p model.Product, i inventory.Inventory)
}
