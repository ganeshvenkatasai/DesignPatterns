package state

import (
	"Vendingmachine/inventory"
	"Vendingmachine/model"
	"errors"
)

type NoCoinInsertedState struct {
}

func (s *NoCoinInsertedState) InsertCoin() error {
	return nil
}

func (s *NoCoinInsertedState) PressButton(amount int, p model.Product, i inventory.Inventory) error {
	return errors.New("Please insert coin")
}

func (s *NoCoinInsertedState) Dispense(p model.Product, i inventory.Inventory) {

}
