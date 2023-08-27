package state

import "Vendingmachine/model"

type State interface {
	insertCoin()
	pressButton()
	dispense() *model.Product
}
