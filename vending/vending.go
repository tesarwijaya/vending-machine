package vending

import (
	"fmt"
	"strconv"
	"strings"
)

var VALID_COIN = []int{10, 50, 100, 500}

// Item ...
type Item struct {
	Price int
	Name  string
	Slot  string
}

// Coin ...
type Coin struct {
	Value int
}

// Vending ...
type Vending struct {
	Balance       int
	Items         map[string][]*Item
	Coins         map[int][]*Coin
	CoinsInserted []*Coin
	ReturnGate    []*Coin
	Outlet        []*Item
}

// InitVending ...
func InitVending() *Vending {
	return &Vending{
		Items: map[string][]*Item{
			"1": []*Item{},
			"2": []*Item{},
			"3": []*Item{},
		},
		Coins: map[int][]*Coin{
			10:  []*Coin{},
			50:  []*Coin{},
			100: []*Coin{},
			500: []*Coin{},
		},
	}
}

// InputMap ...
func (v *Vending) InputMap(input string) {
	inputs := strings.Split(input, " ")
	var cmd, arg string

	cmd = inputs[0]
	if len(inputs) > 1 {
		arg = inputs[1]
	}

	switch cmd {
	case "1":
		value, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Machine error: ", err.Error())
		}

		if err := v.CoinInsert(value); err != nil {
			fmt.Println("Machine error: ", err.Error())
		}
		v.Print()
	case "2":
		if err := v.ItemPurchase(arg); err != nil {
			fmt.Println("Machine error: ", err.Error())
		}

		v.Print()
	case "3":
		if err := v.ItemGet(); err != nil {
			fmt.Println("Machine error: ", err.Error())
		}

		v.Print()
	case "4":
		if err := v.CoinReturnChange(); err != nil {
			fmt.Println("Machine error: ", err.Error())
		}

		v.Print()
	case "5":
		if err := v.CoinGet(); err != nil {
			fmt.Println("Machine error: ", err.Error())
		}

		v.Print()
	default:
		fmt.Println("Machine error: Invalid command")

		v.Print()
	}
}

// Print ...
func (v *Vending) Print() {
	fmt.Println("-------------------------------------------")

	fmt.Println("[Input Amount] ", v.Balance)

	fmt.Println("[Change]")
	fmt.Println("   100 JPY     ", v.IsCoinChangeable(10))
	fmt.Println("   10 JPY      ", v.IsCoinChangeable(10))

	fmt.Println("[Return Gate]")
	for _, i := range v.ReturnGate {
		fmt.Println("   ", i.Value, " JPY")
	}

	fmt.Println("[Items for sales]")
	for k, i := range v.Items {
		if _, text := v.IsItemPurchaseable(k); text != "Hide" {
			fmt.Println("   ", k, i[0].Name, i[0].Price, "JPY", text)
		}
	}

	fmt.Println("[Outlet]")
	for _, i := range v.Outlet {
		fmt.Println("   ", i.Name)
	}

	fmt.Println("-------------------------------------------")

	fmt.Print("Select command: ")
}
