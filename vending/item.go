package vending

import (
	"errors"
)

// IsItemPurchaseable ...
func (v *Vending) IsItemPurchaseable(slot string) (bool, string) {
	if (len(v.Items[slot]) != 0) && (v.Balance >= v.Items[slot][0].Price) {
		return true, "Available"
	}

	if (len(v.Items[slot]) != 0) && (v.Balance < v.Items[slot][0].Price) {
		return false, "Hide"
	}

	return false, "Sold out"
}

// ItemGet ...
func (v *Vending) ItemGet() error {
	if len(v.Outlet) == 0 {
		return errors.New("Outlet empty")
	}

	v.Outlet = []*Item{}

	return nil
}

// ItemPurchase ...
func (v *Vending) ItemPurchase(slot string) error {
	items := v.Items[slot]
	if items == nil {
		return errors.New("Invalid slot")
	}
	item := items[0]

	if canPurchase, _ := v.IsItemPurchaseable(item.Slot); canPurchase {
		v.Balance = v.Balance - item.Price

		v.Items[slot] = v.Items[slot][1:]
		v.Outlet = append(v.Outlet, item)

		return nil
	}

	return errors.New("We are sorry, you cannot purchase this item. Please select available item")

}

// ItemStock ...
func (v *Vending) ItemStock(i *Item) {
	v.Items[i.Slot] = append(v.Items[i.Name], i)
}
