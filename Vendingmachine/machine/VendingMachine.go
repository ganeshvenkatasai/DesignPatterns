package machine

import (
	"Vendingmachine/inventory"
	"Vendingmachine/state"
)

type vendingMachine struct {
	noCoinInsertedState state.State
	coinInsertedState   state.State
	dispensingState     state.State
	machineState        state.State
	amount              int
	inventory           inventory.Inventory
}

func NewVendingMachine() VendingMachine {
	return &vendingMachine{}
}

func (v vendingMachine) AddProduct() {

}
