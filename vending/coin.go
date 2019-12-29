package vending

import (
	"errors"

	"github.com/tesarwijaya/vending-machine/lib"
)

// CoinGet ...
func (v *Vending) CoinGet() error {
	if len(v.ReturnGate) == 0 {
		return errors.New("No coin in return get")
	}

	v.ReturnGate = []*Coin{}
	return nil
}

// CoinInsert ...
func (v *Vending) CoinInsert(value int) error {
	_, valid := lib.FindInSlice(VALID_COIN, value)

	if valid {
		c := Coin{Value: value}

		if (len(v.Coins[10]) == 0) &&
			(c.Value == 50 ||
				c.Value == 100 ||
				c.Value == 500) {
			return errors.New("Sorry we cannot accept this coin for now")
		}

		if (len(v.Coins[100]) == 0) && (c.Value == 500) {
			return errors.New("Sorry we cannot accept this coin for now")
		}

		v.Balance = v.Balance + c.Value

		v.CoinsInserted = append(v.CoinsInserted, &c)
		return nil
	}

	return errors.New("Coin is not valid")
}

// CoinReturnChange ...
func (v *Vending) CoinReturnChange() error {
	if v.Balance == v.CoinInsertedTotal() {
		return errors.New("You haven't purchase anything yet")
	}
	v.CoinInsertedAddToStock()

	amount := v.Balance
	h := (amount / 100)
	amount = amount % 100
	j := (amount / 10)

	v.ReturnGate = append(v.ReturnGate, v.Coins[100][:h]...)
	v.Coins[100] = v.Coins[100][h:]

	v.ReturnGate = append(v.ReturnGate, v.Coins[10][:j]...)
	v.Coins[10] = v.Coins[10][j:]

	v.Balance = 0

	return nil
}

// CoinInsertedAddToStock ...
func (v *Vending) CoinInsertedAddToStock() {
	for _, i := range v.CoinsInserted {
		v.CoinStock(i)
	}
	v.CoinsInserted = []*Coin{}
}

// CoinInsertedTotal ...
func (v *Vending) CoinInsertedTotal() int {
	total := 0

	if len(v.CoinsInserted) != 0 {
		for _, i := range v.CoinsInserted {
			total = total + i.Value
		}
	}

	return total
}

// CoinStock ...
func (v *Vending) CoinStock(c *Coin) {
	v.Coins[c.Value] = append(v.Coins[c.Value], c)
}

// IsCoinChangeable ...
func (v *Vending) IsCoinChangeable(value int) bool {
	if value == 10 {
		return !(len(v.Coins[10]) < 9)
	}

	if value == 100 {
		return !(len(v.Coins[10]) < 4)
	}

	return false
}
