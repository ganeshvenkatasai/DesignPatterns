package machine

import (
	"Vendingmachine/inventory"
	"Vendingmachine/model"
	"Vendingmachine/state"
	"fmt"
	"log"
)

type VendingMachine struct {
	amount       int
	inventory    inventory.Inventory
	machineState state.State
}

func NewVendingMachine(i inventory.Inventory) *VendingMachine {
	return &VendingMachine{
		amount:       0,
		inventory:    i,
		machineState: &state.NoCoinInsertedState{},
	}
}

func (v *VendingMachine) AddProduct(p model.Product) {
	v.inventory.AddProduct(p)
}

func (v *VendingMachine) InsertCoin(coin int) {
	err := v.machineState.InsertCoin()
	if err != nil {
		log.Println(err)
		return
	}
	v.amount += coin
	v.SetCoinInsertedState()
}

func (v *VendingMachine) PressButton(p model.Product) {
	err := v.machineState.PressButton(v.amount, p, v.inventory)
	if err != nil {
		log.Println(err)
		return
	}
	v.SetDispenseState()
	v.machineState.Dispense(p, v.inventory)
	fmt.Println("Please collect your", p.Name, "along with change Rs.", v.amount-p.Price)
	v.SetNoCoinInsertedState()
}

func (v *VendingMachine) SetNoCoinInsertedState() {
	v.machineState = v.GetNoCoinInsertedState()
}

func (v *VendingMachine) GetNoCoinInsertedState() state.State {
	return &state.NoCoinInsertedState{}
}

func (v *VendingMachine) SetCoinInsertedState() {
	v.machineState = v.GetCoinInsertedState()
}

func (v *VendingMachine) GetCoinInsertedState() state.State {
	return &state.CoinInsertedState{}
}

func (v *VendingMachine) SetDispenseState() {
	v.machineState = v.GetDispenseState()
}

func (v *VendingMachine) GetDispenseState() state.State {
	return &state.DispensingState{}
}
