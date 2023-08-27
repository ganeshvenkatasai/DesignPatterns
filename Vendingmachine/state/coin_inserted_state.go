package state

import (
	"Vendingmachine/inventory"
	"Vendingmachine/model"
	"errors"
)

type CoinInsertedState struct {
}

func (state CoinInsertedState) InsertCoin() error {
	return nil
}

func (state CoinInsertedState) PressButton(amount int, p model.Product, i inventory.Inventory) error {
	if val, ok := i.ProductCountMap[p]; ok {
		if val > 0 {
			i.ProductCountMap[p]--
			return nil
		} else {
			return errors.New("Product unavailable")
		}
	} else {
		return errors.New("No such product available")
	}
}

func (state CoinInsertedState) Dispense(p model.Product, i inventory.Inventory) {

}
