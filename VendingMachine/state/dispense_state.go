package state

import (
	"VendingMachine/inventory"
	"VendingMachine/model"
	"errors"
)

type DispensingState struct {
}

func (state DispensingState) InsertCoin() error {
	return errors.New("Please insert coin after dispensing")
}

func (state DispensingState) PressButton(amount int, p model.Product, i inventory.Inventory) error {
	return nil
}

func (state DispensingState) Dispense(p model.Product, i inventory.Inventory) {
	i.ProductCountMap[p]--
}
