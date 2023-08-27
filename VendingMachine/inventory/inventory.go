package inventory

import "VendingMachine/model"

type Inventory struct {
	ProductCountMap map[model.Product]int
}

func NewInventory() *Inventory {
	return &Inventory{
		ProductCountMap: make(map[model.Product]int, 0),
	}
}

func (i *Inventory) AddProduct(p model.Product) {
	i.ProductCountMap[p]++
}
