package state

import "Vendingmachine/model"

type NoCoinInsertedState struct {
}

type CoinInsertedState struct {
}

type DispensingState struct {
}

func (state NoCoinInsertedState) insertCoin() {
	
}

func (state NoCoinInsertedState) pressButton() {

}

func (state NoCoinInsertedState) dispense() *model.Product {
	return nil
}

func (state CoinInsertedState) insertCoin() {

}

func (state CoinInsertedState) pressButton() {

}

func (state CoinInsertedState) dispense() *model.Product {
	return nil
}

func (state DispensingState) insertCoin() {

}

func (state DispensingState) pressButton() {

}

func (state DispensingState) dispense() *model.Product {
	return nil
}
